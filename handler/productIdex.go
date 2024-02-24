package handler

import (
	"net/http"

	"github.com/br-jeff/simple-go-htmx/repository"
)

func ProductIndex(w http.ResponseWriter) {
	products := repository.FindAll()
	templates.ExecuteTemplate(w, "Index", products)
}
