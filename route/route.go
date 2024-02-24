package route

import (
	"net/http"

	"github.com/br-jeff/simple-go-htmx/controller"
)

func Bootstrap() {
	http.HandleFunc("/", controller.ProductIndex)
	http.HandleFunc("/new", controller.ProductNew)

}
