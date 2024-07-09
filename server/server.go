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

var server = NewServer()
var addr = flag.String("addr", "pjalv.com", "http service address")

func statusChecker() {
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
					server.mu.Lock()
					upService = localURL
					server.mu.Unlock()
				} else {
					log.Printf("Domain %s is up. Status code: %d", domains, resp.StatusCode)
					for key, val := range domainList {
						if val == domain {
							if upService != key {
								server.mu.Lock()
								upService = key
								server.mu.Unlock()
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
func WSHandler(w http.ResponseWriter, r *http.Request) {
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
	server.mu.Lock()
	server.conns[ws] = &ConnectionInfo{
		Agent: agentString}
	server.mu.Unlock()
	defer func() {
		if server.conns[ws].Agent == "commander" {
			go commandSender(channelCommand)
		}
		ws.Close()
		delete(server.conns, ws)
	}()
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			if server.conns[ws].Agent == "commander" {
				log.Print("Commander WS down, restarting")
			}
			return
		}
		log.Printf("Received message from [%s]: %s", server.conns[ws].Agent, string(message))

		broadcast(messageType, message, ws)
	}
}
func broadcast(messageType int, b []byte, authWS *websocket.Conn) {
	log.Printf("BROADCASTING FROM %s", server.conns[authWS].Agent)
	if server.conns[authWS].Agent == "client" || server.conns[authWS].Agent == "commander" {
		fmt.Println("Sending to broker...")
		for ws := range server.conns {
			if server.conns[ws].Agent == "broker" {
				go func(ws *websocket.Conn) {
					if err := ws.WriteMessage(messageType, b); err != nil {
						fmt.Println("Write Error: ", err)
					}
				}(ws)
			}
		}
	}
	if server.conns[authWS].Agent == "broker" {
		for ws := range server.conns {
			if server.conns[ws].Agent == "client" {
				go func(ws *websocket.Conn) {
					if err := ws.WriteMessage(messageType, b); err != nil {
						fmt.Println("Write Error: ", err)
					}
				}(ws)
			}
		}
	}
}

func commandSender(ch chan []byte) {
	fmt.Println("Starting command sender")
	key := []byte(os.Getenv("JWT_SECRET"))
	t := jwt.New(jwt.SigningMethodHS256)
	s, err := t.SignedString(key)
	if err != nil {
		return
	}
	u := url.URL{Scheme: "wss", Host: *addr, Path: "/ws", RawQuery: "token=" + s + "&agent=commander"}
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

func CommandHandler(w http.ResponseWriter, r *http.Request) {
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
	go postCommand(command, channelCommand)
	w.Write([]byte("Processing Request!"))
	fmt.Println("Received command:", command)
}

func postCommand(command string, ch chan []byte) {

	log.Println("Service:", upService)
	var headers map[string]string
	var url string
	switch string(upService) {
	case "openai":
		headers = map[string]string{"Authorization": "Bearer " + os.Getenv("OPENAI_APIKEY"), "Content-Type": "application/json"}
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
				return "gpt-4o"
			default: // default is local
				return "gpt-3.5-turbo"
			}
		}(upService),
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": "You are an assistant tasked with controlling two IoT devices: a fan and an RGB LED strip. You will perform tasks related to these devices and respond with an appropriate JSON object  ALL OBJECT KEYS and STRINGS MUST HAVE DOUBLE QUOTES. If the input is not recognizable, respond with an error object: { response: 'error' }.  Fan Device Schema: Power Control: If asked to turn the fan on or off, respond with: { response: 'ok', topic: 'fan/control', payload_format: 'INT', payload: '{0 for OFF or 1 for ON}' }.  Speed Control: If asked to set the fan speed, respond with a number between 96 and 1024 for the duty cycle: { response: 'ok', topic: 'fan/control', payload_format: 'INT', payload: 'number between 96 and 1024' }.  Function Control: If asked to set the Fan to 'breeze' mode, then respond with this object : {response: 'ok', topic: 'fan/control', payload_format: 'JSON', payload: {function: 1}} LED Strip: Power Control: If asked to turn the LEDs on or off, respond with: { response: 'ok', topic: 'led/control/power', payload_format: 'INT', payload: '{0 for off, 1 for on}' }.  Color Control: If asked to change the color of the LEDs, respond with the RGB values of the color: { response: 'ok', topic: 'led/control/color', payload_format: 'JSON', payload: { red: {R value}, green: {G value}, blue: {B value} } }.  Function Control: If asked to set the strip to 'static rainbow' mode, then respond with this object : {response: 'ok', topic: 'led/control/color', payload_format: 'JSON', payload: {function: 1}} If asked to set the strip to 'Trailing rainbow' mode, then respond with this object : {response: 'ok', topic: 'led/control/color', payload_format: 'JSON', payload: {function: 2}}",
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
		go postCommand(command, ch)
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
		go postCommand(command, ch)
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
