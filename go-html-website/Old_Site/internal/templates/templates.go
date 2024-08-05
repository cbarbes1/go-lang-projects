package templates

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "path/filepath"
)

var templates *template.Template

func LoadTemplates() {
    pattern := filepath.Join("web", "templates", "**", "*.html")
    log.Println("Loading templates from:", pattern)
    
    // List all files matching the pattern
    matches, err := filepath.Glob(pattern)
    if err != nil {
        log.Fatal("Error finding template files:", err)
    }
    log.Println("Found template files:", matches)

    var err error
    templates, err = template.ParseGlob(pattern)
    if err != nil {
        log.Fatal("Error parsing templates:", err)
    }

    // List all defined templates
    for _, t := range templates.Templates() {
        log.Println("Loaded template:", t.Name())
    }

    log.Println("Templates loaded successfully")
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    if templates == nil {
        log.Println("Templates are not loaded")
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    t := templates.Lookup(tmpl)
    if t == nil {
        log.Printf("Template %s not found", tmpl)
        http.Error(w, "Template Not Found", http.StatusInternalServerError)
        return
    }

    err := t.Execute(w, data)
    if err != nil {
        log.Println("Error executing template:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}