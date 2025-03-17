package main

import (
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"
	"website/server"
	components "website/templates"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	_ = r
	http.Error(w, "404 page not found", http.StatusNotFound)
}

func redirectToURL(w http.ResponseWriter, r *http.Request, URL string) {
	http.Redirect(w, r, URL, http.StatusSeeOther)
}

func setupServer(secure bool, router *chi.Mux) {
	// Start the HTTPS server
	if secure {
		err := http.ListenAndServe(":3000", router)
		if err != nil {
			panic(err)
		}
	} else {
		err := http.ListenAndServe(":8080", router)
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
	sdp := server.NewServer()
	r := chi.NewRouter()
	os.Setenv("TZ", "America/Los_Angeles")
	time.LoadLocation("America/Los_Angeles")

	if err := godotenv.Load(".env"); err != nil {
		log.Print("Error loading .env file")
	}
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to Favicon")
		http.ServeFile(w, r, "./static/favicon.ico")
	})
	r.Get("/file/*", fileHandler)
	r.Get("/ws", sdp.WSHandler)
	r.Post("/commands", sdp.CommandHandler)
	r.Get("/chipotle-bot", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to Chipotle Bot")
		redirectToURL(w, r, "https://github.com/PJalv/chipbot-nba-finals-2023")
	})
	r.Get("/jorge", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to Jorge")
		redirectToURL(w, r, "https://www.linkedin.com/in/jorgelsuarez")
	})
	r.Get("/contact", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to Contact")
		Contact := components.Post{
			Title:       "Contact",
			Description: "Jorge Luis Suarez",
			Content:     "",
			RawTitle:    "Contact",
		}

		components.Header(Contact).Render(r.Context(), w)
		components.NavBar().Render(r.Context(), w)
		components.Contact().Render(r.Context(), w)
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
	r.Get("/nixos", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Req to NIXOS")
		redirectToURL(w, r, "https://pjalv.com/blog/from-arch-to-nixos-a-journey-into-declarative-system-configuration")
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
		temp := make([]components.Post, len(components.Posts))
		copy(temp, components.Posts)
		slices.Reverse(temp)
		Blog := components.Post{
			Title:       "Blog",
			Description: "Jorge Luis Suarez",
			Content:     "",
			RawTitle:    "Blog",
		}
		components.Header(Blog).Render(r.Context(), w)
		log.Println(len(temp))
		for _, data := range temp {
			log.Println(data.Title)
		}
		components.BlogIndex(temp).Render(r.Context(), w)
	})
	r.Post("/blog-data-rev", func(w http.ResponseWriter, r *http.Request) {
		var rev bool
		reverse := r.Header.Get("reverse")
		if reverse == "true" {
			rev = true
		} else {
			rev = false
		}
		temp := make([]components.Post, len(components.Posts))
		copy(temp, components.Posts)
		if rev {
			slices.Reverse(temp)
		}
		components.CompBlogData(temp, rev).Render(r.Context(), w)
	})
	r.Get("/blog-update", func(w http.ResponseWriter, r *http.Request) {
		components.MDConvert()
	})
	r.Get("/interviews", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to Interviews")
		Interviews := components.Post{
			Title:       "Interviews",
			Description: "Jorge Luis Suarez - Interview Collection",
			Content:     "",
			RawTitle:    "Interviews",
		}

		components.Header(Interviews).Render(r.Context(), w)
		components.NavBar().Render(r.Context(), w)
		components.Interviews().Render(r.Context(), w)
	})

	r.Get("/blog/*", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		log.Println(path)
		parts := strings.Split(path, "/")
		if len(parts) < 3 {
			http.NotFound(w, r)
			return
		}
		var id int = -99
		for index, post := range components.Posts {
			if strings.ToLower(post.Title) == parts[2] {
				id = index
			}
		}
		if id == -99 {
			redirectToURL(w, r, "https://pjalv.com")
		} else {
			Page := components.Posts[id]
			Page.Title = strings.ReplaceAll(components.Posts[id].Title, "-", " ")
			components.Header(Page).Render(r.Context(), w)
			components.BlogPage(components.Posts[id]).Render(r.Context(), w)
		}
	})

	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to Home")
		go sendDiscWebhook(r)
		redirectToURL(w, r, "https://pjalv.com")
	})

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		redirectToURL(w, r, "https://pjalv.com")
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		// Check if the request is using www
		if host == "www.pjalv.com" {
			log.Println("GOT WWW REQ")
			// Redirect to non-www
			http.Redirect(w, r, "https://pjalv.com"+r.URL.String(), http.StatusMovedPermanently)
			return
		}
		Landing := components.Post{
			Title:       "PJalv",
			Description: "Jorge Luis Suarez",
			Content:     "",
			RawTitle:    "Jorge Luis Suarez",
		}
		components.Header(Landing).Render(r.Context(), w)
		components.NavBar().Render(r.Context(), w)
		components.NewestBlogPost(components.Posts).Render(r.Context(), w)
		components.Landing().Render(r.Context(), w)
	})
	// go commandSender(channelCommand)
	// go statusChecker()
	components.MDConvert()
	components.InterviewsConvert()
	args := os.Args
	if args[1] == "true" {
		setupServer(true, r)
	} else {
		setupServer(false, r)
	}
}
func sendDiscWebhook(r *http.Request) {
	webhookURL := "https://discord.com/api/webhooks/1288971178752479303/634SY6yuZkUayFdt-mBFK-HW-JU-ddvMAzAKXGCXz2rYbdrYxlIN0ewYlTiD6fjcpjhi" // Replace with your actual webhook URL
	ip := r.RemoteAddr
	message := `{"content": "Home page clicked! IP: ` + ip + `"}`
	_, err := http.Post(webhookURL, "application/json", strings.NewReader(message))
	if err != nil {
		log.Println("Error sending Discord notification:", err)
	}
}
