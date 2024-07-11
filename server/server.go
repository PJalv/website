package server

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type ServerInterface interface {
	WSHandler(w http.ResponseWriter, r *http.Request)
	CommandHandler(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	conns map[*websocket.Conn]*ConnectionInfo
	mu    sync.Mutex
}
type ConnectionInfo struct {
	Agent string
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]*ConnectionInfo),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var localURL = "https://associated-walked-gore-wise.trycloudflare.com"
var upService = "openai"

var channelCommand = make(chan []byte)

var addr = flag.String("addr", "pjalv.com", "http service address")

func (s *Server) statusChecker() {
	domainList := map[string]string{"local": localURL}
	for {
		for _, domain := range domainList {
			go func(domains string) {
				log.Println("Checking domain:", domains)
				client := &http.Client{
					Timeout: 3 * time.Second, // Set timeout to 10 seconds
				}
				// Make a GET request with the custom client
				ctx, cancel := context.WithTimeout(context.Background(), client.Timeout)
				defer cancel()
				req, err := http.NewRequestWithContext(ctx, "PATCH", domains, nil)
				if err != nil {
					fmt.Println("Error creating request:", err)
					return
				}
				resp, err := client.Do(req)
				if err != nil {
					if ctx.Err() == context.DeadlineExceeded {
						log.Printf("Timeout when checking domain %s", domains)
					}
					return
				}
				log.Println(resp.StatusCode)
				if resp.StatusCode != 421 && resp.StatusCode != 501 {
					log.Printf("Domain %s is down. Status code: %d", domains, resp.StatusCode)
					s.mu.Lock()
					upService = localURL
					s.mu.Unlock()
				} else {
					log.Printf("Domain %s is up. Status code: %d", domains, resp.StatusCode)
					for key, val := range domainList {
						if val == domain {
							if upService != key {
								s.mu.Lock()
								upService = key
								s.mu.Unlock()
								continue
							}
						}
					}

				}
			}(domain)
			time.Sleep(7 * time.Second)
		}
	}
}

func verifyToken(tokenString string, secretKey []byte) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	// fmt.Print(token)
	return nil
}

func (s *Server) WSHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
		fmt.Println("Token not provided")
		http.Error(w, "Token not provided", http.StatusBadRequest)
		return
	}
	// Verify the token
	if err := verifyToken(tokenString, []byte(os.Getenv("JWT_SECRET"))); err != nil {
		fmt.Println("Token verification failed:", err)
		http.Error(w, "Token verification failed", http.StatusUnauthorized)
		return
	}
	agentString := r.URL.Query().Get("agent")
	if agentString == "" {
		fmt.Println("Agent not provided")
		http.Error(w, "Agent not provided", http.StatusBadRequest)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {

		log.Println(err)
		return
	}
	fmt.Println("New incoming connection from Client:", r.RemoteAddr)
	fmt.Printf("URL: %v\n", r.URL)
	s.mu.Lock()
	s.conns[ws] = &ConnectionInfo{
		Agent: agentString}
	s.mu.Unlock()
	defer func() {
		if s.conns[ws].Agent == "commander" {
			go s.commandSender(channelCommand)
		}
		ws.Close()
		delete(s.conns, ws)
	}()
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			if s.conns[ws].Agent == "commander" {
				log.Print("Commander WS down, restarting")
			}
			return
		}
		log.Printf("Received message from [%s]: %s", s.conns[ws].Agent, string(message))

		s.broadcast(messageType, message, ws)
	}
}
func (s *Server) broadcast(messageType int, b []byte, authWS *websocket.Conn) {
	log.Printf("BROADCASTING FROM %s", s.conns[authWS].Agent)
	if s.conns[authWS].Agent == "client" || s.conns[authWS].Agent == "commander" {
		fmt.Println("Sending to broker...")
		for ws := range s.conns {
			if s.conns[ws].Agent == "broker" {
				go func(ws *websocket.Conn) {
					if err := ws.WriteMessage(messageType, b); err != nil {
						fmt.Println("Write Error: ", err)
					}
				}(ws)
			}
		}
	}
	if s.conns[authWS].Agent == "broker" {
		for ws := range s.conns {
			if s.conns[ws].Agent == "client" {
				go func(ws *websocket.Conn) {
					if err := ws.WriteMessage(messageType, b); err != nil {
						fmt.Println("Write Error: ", err)
					}
				}(ws)
			}
		}
	}
}

func (s *Server) commandSender(ch chan []byte) {
	fmt.Println("Starting command sender")
	key := []byte(os.Getenv("JWT_SECRET"))
	t := jwt.New(jwt.SigningMethodHS256)
	signedString, err := t.SignedString(key)
	if err != nil {
		return
	}
	u := url.URL{Scheme: "wss", Host: *addr, Path: "/ws", RawQuery: fmt.Sprintf("token=%s&agent=commander", signedString)}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
		return
	}
	defer func() {
		log.Print("COMMANDER CONNECTION LOST... ")
	}()
	for {
		select {
		case payload, ok := <-ch:
			if !ok {
				fmt.Print("Channel closed")
				return
			}
			if payload != nil {
				log.Printf("Payload from channel: %s", payload)
			}
			err := c.WriteMessage(websocket.TextMessage, payload)
			if err != nil {
				log.Printf("Error writing to WS server: %s", err)
				ch <- payload
				return
			}

		default:
			continue
		}
	}
}

func (s *Server) CommandHandler(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	command, ok := params["command"]
	if !ok {
		fmt.Println("command key not found in json request")
		return
	}
	go s.postCommand(command, channelCommand)
	w.Write([]byte("Processing Request!"))
	fmt.Println("Received command:", command)
}

func (s *Server) postCommand(command string, ch chan []byte) {
	log.Println("received command", command)
	log.Println("Service:", upService)
	var headers map[string]string
	var url string
	switch string(upService) {
	case "openai":
		headers = map[string]string{"Authorization": "Bearer " + os.Getenv("OPENAI_API_KEY"), "Content-Type": "application/json"}
		url = "https://api.openai.com/v1/chat/completions"
	default:
		headers = map[string]string{}
		url = localURL + "/v1/chat/completions"
	}
	var body interface{} = map[string]interface{}{
		"model": func(service string) string {
			switch service {
			case "local":
				return "koboldcpp/Noromaid-v0.4-Mixtral-Instruct-8x7b.q3_k_m"
			case "openai":
				return os.Getenv("OPENAI_MODEL")
			default: // default is local
				return "gpt-3.5-turbo"
			}
		}(upService),
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": os.Getenv("IOT_COMMAND"),
			},
			{"role": "user", "content": command},
		},
	}
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return
	}
	reader := bytes.NewReader(jsonStr)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Print("Error in post request.\n")
		return
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	log.Println("doing request")
	resp, err := client.Do(req)
	if err != nil {
		log.Print("Error in post request. Falling back to OpenAI...\n")
		upService = "openai"
		go s.postCommand(command, ch)
		return
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print("Error in post request.\n")
		return
	}
	fmt.Printf("response status code: %d\n", resp.StatusCode)
	if resp.StatusCode != 200 {
		log.Println("Request failed, falling back to OpenAI...")
		upService = "openai"
		go s.postCommand(command, ch)
		return
	}
	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return
	}
	strResponse := response.Choices[0].Message.Content
	fmt.Println(strResponse)

	ch <- []byte(strResponse)
}

type Response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}
