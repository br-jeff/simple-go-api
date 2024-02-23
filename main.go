package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	connection := "user=usernamepg dbname=db password=password123 host=postgresgo sslmode=disable"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Id          int
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
	db := connectDB()

	selectProduct, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64
		var createdAt time.Time

		err = selectProduct.Scan(&id, &name, &description, &price, &quantity, &createdAt)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	templat.ExecuteTemplate(w, "Index", products)
	defer db.Close() // Close connection
}
