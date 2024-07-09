package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"
	"website/server"
	"website/templates"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type Post struct {
	Date    time.Time
	Title   string
	Content string
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 page not found", http.StatusNotFound)
}

func redirectToURL(w http.ResponseWriter, r *http.Request, URL string) {
	http.Redirect(w, r, URL, http.StatusSeeOther)
}

func setupServer(secure bool, router *chi.Mux) {
	// Start the HTTPS server
	if secure {
		certFile := "/etc/letsencrypt/live/pjalv.com/fullchain.pem"
		keyFile := "/etc/letsencrypt/live/pjalv.com/privkey.pem"
		// Load the TLS certificate and key
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			panic(err)
		}
		tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
		server := &http.Server{
			Addr:      ":443",
			TLSConfig: tlsConfig,
			Handler:   router,
		}
		err = server.ListenAndServeTLS("", "")
		if err != nil {
			panic(err)
		}
	} else {
		err := http.ListenAndServe(":3000", router)
		if err != nil {
			panic(err)
		}
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the requested file path from the URL
	filePath := "./static/" + r.URL.Path[len("/file/"):]

	// Check if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, return 404 Not Found
			http.NotFound(w, r)
			return
		}
		// For other errors (e.g., permission denied), return 500 Internal Server Error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, filePath)
}
func main() {
	r := chi.NewRouter()
	os.Setenv("TZ", "America/Los_Angeles")
	time.LoadLocation("America/Los_Angeles")

	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("Error loading .env file")
	}
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to Favicon")
		http.ServeFile(w, r, "./static/favicon.ico")
	})
	r.Get("/file/*", fileHandler)
	r.Get("/ws", server.WSHandler)
	r.Post("/commands", server.CommandHandler)
	r.Get("/chipotle-bot", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to Chipotle Bot")
		redirectToURL(w, r, "https://github.com/PJalv/chipbot-nba-finals-2023")
	})
	r.Get("/jorge", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to Jorge")
		redirectToURL(w, r, "https://www.linkedin.com/in/jorgelsuarez")
	})
	r.Get("/iot", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to GitHub Repo")
		redirectToURL(w, r, "https://github.com/PJalv/SDP-IoT-Device-Control")
	})
	r.Get("/iot-presentation", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to Video Presentation")
		redirectToURL(w, r, "https://www.youtube.com/watch?v=jLNg5vOTuZE")
	})
	r.Get("/voicedemo", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Req to YT VIDEO")
		redirectToURL(w, r, "https://www.youtube.com/watch?v=W1N3hpLCcY0")
	})
	r.Get("/resume", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Req to Resume")
		redirectToURL(w, r, "https://docs.google.com/document/d/1ynXWoHiHe3NNCNCZin4Slt91ttA6x-v9jrYs64DE5KU/edit?usp=sharing")
	})
	r.Get("/dylan", func(w http.ResponseWriter, r *http.Request) {
		log.Println("REQ to DYLAN")
		redirectToURL(w, r, "https://www.linkedin.com/in/dylanstlaurent")
	})
	r.Get("/blog", func(w http.ResponseWriter, r *http.Request) {
		// templates.MDConvert()
		templates.Header("Blog - PJalv").Render(r.Context(), w)
		log.Println(len(templates.Posts))
		templates.BlogIndex(templates.Posts).Render(r.Context(), w)
	})
	r.Post("/blog-data-rev", func(w http.ResponseWriter, r *http.Request) {
		var rev bool
		reverse := r.Header.Get("reverse")
		if reverse == "true" {
			rev = true
		} else {
			rev = false
		}
		temp := make([]templates.Post, len(templates.Posts))
		copy(temp, templates.Posts)
		if rev {
			slices.Reverse(temp)
		}
		templates.CompBlogData(temp, rev).Render(r.Context(), w)
	})
	r.Get("/blog-update", func(w http.ResponseWriter, r *http.Request) {
		templates.MDConvert()
	})
	r.Get("/blog/*", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		log.Println(path)
		// Example: /blog/123 -> ["", "blog", "123"]
		parts := strings.Split(path, "/")
		if len(parts) < 3 {
			http.NotFound(w, r)
			return
		}
		var id int = -99
		for index, post := range templates.Posts {
			if strings.ToLower(post.Title) == parts[2] {
				id = index
			}
		}
		if id == -99 {
			redirectToURL(w, r, "https://pjalv.com")
		} else {
			templates.Header(strings.ReplaceAll(templates.Posts[id].Title, "-", " ")+" - PJalv").Render(r.Context(), w)
			templates.BlogPage(templates.Posts[id]).Render(r.Context(), w)
		}
	})
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		templates.Header("Jorge Luis Suarez - PJalv").Render(r.Context(), w)
		templates.NavBar().Render(r.Context(), w)
		templates.NewestBlogPost(templates.Posts).Render(r.Context(), w)
		templates.Landing().Render(r.Context(), w)
	})
	// go commandSender(channelCommand)
	// go statusChecker()
	templates.MDConvert()
	args := os.Args
	if args[1] == "true" {
		setupServer(true, r)
	} else {
		setupServer(false, r)
	}
}
