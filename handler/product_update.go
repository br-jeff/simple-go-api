package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/br-jeff/simple-go-htmx/repository"
)

func ProductUpdate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		log.Println("erro when try to convert id", err)
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)

	if err != nil {
		log.Println("erro when try to convert price", err)
	}

	quantity, err := strconv.Atoi(r.FormValue("quantity"))

	if err != nil {
		log.Println("erro when try to convert quantity", err)
	}

	repository.ProductUpdate(id, name, description, quantity, price)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
