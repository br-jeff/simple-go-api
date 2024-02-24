package main

import (
	"net/http"

	"github.com/br-jeff/simple-go-htmx/route"
)

func main() {
	route.Bootstrap()
	http.ListenAndServe(":8080", nil)
}
