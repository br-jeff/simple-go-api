package handler

import (
	"net/http"

	"github.com/br-jeff/simple-go-htmx/repository"
)

func ProductEdit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := repository.ProductFindById(productId)
	templates.ExecuteTemplate(w, "Edit", product)
}
