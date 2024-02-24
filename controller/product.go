package controller

import (
	"net/http"
	"text/template"

	"github.com/br-jeff/simple-go-htmx/repository"
)

var templat = template.Must(template.ParseGlob("templates/*.html"))

func ProductIndex(w http.ResponseWriter, r *http.Request) {
	products := repository.FindAll()
	templat.ExecuteTemplate(w, "Index", products)
}

func ProductNew(w http.ResponseWriter, r *http.Request) {
	templat.ExecuteTemplate(w, "NewProduct", nil)
}
