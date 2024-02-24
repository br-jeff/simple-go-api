package controller

import (
	"net/http"
	"text/template"

	"github.com/br-jeff/simple-go-htmx/repository"
)

func ProductIndex(w http.ResponseWriter, r *http.Request) {
	var templat = template.Must(template.ParseGlob("templates/*.html"))
	products := repository.FindAll()
	templat.ExecuteTemplate(w, "Index", products)
}
