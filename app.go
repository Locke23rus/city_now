package main

import (
	"html/template"
	"net/http"
)

const (
	publicPath       = "public"
	rootTemplatePath = "templates/index.html"
)

var rootTemplate *template.Template

func handler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, nil)
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, publicPath+r.URL.Path)
}

func init() {
	rootTemplate = template.Must(template.ParseFiles(rootTemplatePath))
}

func main() {
	http.HandleFunc("/css", handleStatic)
	http.HandleFunc("/js", handleStatic)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
