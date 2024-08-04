package templates

import (
    "html/template"
    "log"
    "net/http"
)

var templates *template.Template

func LoadTemplates() {
    var err error
    templates, err = template.ParseGlob("web/templates/**/*.html")
    if err != nil {
        log.Fatal(err)
    }
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
