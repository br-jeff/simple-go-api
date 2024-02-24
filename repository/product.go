package repository

import (
	"time"

	"github.com/br-jeff/simple-go-htmx/config"
	"github.com/br-jeff/simple-go-htmx/model"
)

func FindAll() []model.Product {

	db := config.ConnectDB()

	selectProduct, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := model.Product{}
	products := []model.Product{}

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

	defer db.Close() // Close connection

	return products
}

func InsertProduct(name string, description string, price float64, quantity int) {
	db := config.ConnectDB()

	productInsetTemplate, err := db.Prepare("INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4);")

	if err != nil {
		panic(err.Error())
	}

	productInsetTemplate.Exec(name, description, price, quantity)

	defer db.Close()
}
