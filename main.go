package main

import (
	"html/template"
	"net/http"
)

var templat = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	templat.ExecuteTemplate(w, "Index", nil)
}
