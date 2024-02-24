package controller

import (
	"log"
	"net/http"
	"strconv"
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

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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
}
