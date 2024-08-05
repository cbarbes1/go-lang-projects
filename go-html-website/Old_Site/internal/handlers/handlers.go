package handlers

import (
	"go-htmx-website/internal/templates"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "World"
	}
	response := struct{ Message string }{Message: "Hello, " + name + "!"}
	templates.ExecuteTemplate(w, "hello.html", response)
}
