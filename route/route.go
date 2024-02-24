package route

import (
	"net/http"

	"github.com/br-jeff/simple-go-htmx/controller"
)

func Bootstrap() {
	http.HandleFunc("/", controller.ProductIndex)
	http.HandleFunc("/new", controller.ProductNew)
	http.HandleFunc("/insert", controller.ProductCreate)
	http.HandleFunc("/delete", controller.ProductDelete)
	http.HandleFunc("/edit", controller.ProductEdit)
	http.HandleFunc("/update", controller.ProductUpdate)
}
