package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
    "log"
)

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/hello.html"))

func main() {
    r := mux.NewRouter()
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    r.HandleFunc("/", homeHandler)
    r.HandleFunc("/hello", helloHandler).Methods("POST")

    http.Handle("/", r)
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    if name == "" {
        name = "World"
    }
    response := struct{ Message string }{Message: "Hello, " + name + "!"}
    templates.ExecuteTemplate(w, "hello.html", response)
}

