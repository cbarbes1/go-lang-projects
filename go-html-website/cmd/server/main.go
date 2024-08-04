package main

import (
	"log"
	"net/http"
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
		log.Fatal(err)
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
				log.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add("templates")
	if err != nil {
		log.Fatal(err)
	}

	// Set up the router and handlers
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/hello", handlers.HelloHandler).Methods("POST")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
