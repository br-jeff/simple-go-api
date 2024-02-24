package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/br-jeff/simple-go-htmx/handler"
	"github.com/br-jeff/simple-go-htmx/repository"
)

var templat = template.Must(template.ParseGlob("templates/*.html"))

func ProductIndex(w http.ResponseWriter, r *http.Request) {
	handler.ProductIndex(w)
}

func ProductNew(w http.ResponseWriter, r *http.Request) {
	templat.ExecuteTemplate(w, "NewProduct", nil)
}

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handler.ProductCreate(w, r)
	}
}

func ProductDelete(w http.ResponseWriter, r *http.Request) {
	handler.ProductDelete(w, r)
}

func ProductEdit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := repository.ProductFindById(productId)
	templat.ExecuteTemplate(w, "Edit", product)
}

func ProductUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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
}
