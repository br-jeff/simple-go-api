package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/br-jeff/simple-go-htmx/repository"
)

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")

	priceFloat, err := strconv.ParseFloat(price, 64)

	if err != nil {
		log.Println("Error while convert price")
	}

	quantityInt, err := strconv.Atoi(quantity)

	if err != nil {
		log.Println("Error while convert quantity")
	}

	repository.InsertProduct(name, description, priceFloat, quantityInt)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
