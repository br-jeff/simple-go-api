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
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)

	if err != nil {
		log.Println("Error while convert price")
	}

	quantity, err := strconv.Atoi(r.FormValue("quantity"))

	if err != nil {
		log.Println("Error while convert quantity")
	}

	repository.InsertProduct(name, description, price, quantity)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
