package main

import (
	"net/http"

	"github.com/br-jeff/simple-go-htmx/controller"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	controller.ProductIndex(w, req)

}
