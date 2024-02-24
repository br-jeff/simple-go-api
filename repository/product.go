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

func ProductDelete(id string) {
	db := config.ConnectDB()

	deleteTemplate, err := db.Prepare("DELETE FROM products where id = $1;")

	if err != nil {
		panic(err.Error())
	}

	deleteTemplate.Exec(id)

	defer db.Close()
}

func ProductFindById(id string) model.Product {
	db := config.ConnectDB()

	productQuery, err := db.Query("select * from products where id = $1 ", id)

	if err != nil {
		panic(err.Error())
	}

	productData := model.Product{}

	for productQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64
		var createdAt time.Time

		err = productQuery.Scan(&id, &name, &description, &price, &quantity, &createdAt)

		if err != nil {
			panic(err.Error())
		}

		productData.Id = id
		productData.Name = name
		productData.Description = description
		productData.Price = price
		productData.Quantity = quantity
	}

	defer db.Close()
	return productData
}

func ProductUpdate(id int, name string, description string, quantity int, price float64) {
	db := config.ConnectDB()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id = $5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()
}
