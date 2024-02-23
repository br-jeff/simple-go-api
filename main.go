package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templat = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Iphone 99", Description: "New Smart Phone", Price: 1200.00, Quantity: 12},
		{Name: "Notebook intel 99X", Description: "A New intel notebook", Price: 3000.14, Quantity: 5},
		{"Phone razer", "The best phone", 123.23, 5},
		{"Charger Gorila", "99W fast charger", 134.53, 19},
	}
	templat.ExecuteTemplate(w, "Index", products)
}
