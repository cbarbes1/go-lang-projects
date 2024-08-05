package main

import (
	"log"
	"net/http"
	"path/filepath"
	"go-htmx-website/internal/handlers"
	"go-htmx-website/internal/templates"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
)

func main() {
	// Initial load of templates
	templates.LoadTemplates()

	// Set up the file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Error creating watcher:", err)
	}
	defer watcher.Close()

	// Watch the templates directory for changes
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Template file modified:", event.Name)
					templates.LoadTemplates()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher error:", err)
			}
		}
	}()

	// Add the templates directory to the watcher
	templatesDir := filepath.Join("web", "templates")
	log.Println("Adding watcher for directory:", templatesDir)
	err = watcher.Add(templatesDir)
	if err != nil {
		log.Fatal("Error adding watcher for templates directory:", err)
	}

	// Set up the router and handlers
	r := mux.NewRouter()
	
	// Serve static files
	staticDir := filepath.Join("web", "static")
	log.Println("Serving static files from:", staticDir)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/hello", handlers.HelloHandler).Methods("POST")

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}