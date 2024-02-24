package handler

import (
	"net/http"

	"github.com/br-jeff/simple-go-htmx/repository"
)

func ProductDelete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	repository.ProductDelete(productId)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
