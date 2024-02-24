package controller

import (
	"net/http"

	"github.com/br-jeff/simple-go-htmx/handler"
)

func ProductIndex(w http.ResponseWriter, r *http.Request) {
	handler.ProductIndex(w)
}

func ProductNew(w http.ResponseWriter, r *http.Request) {
	handler.ProductNew(w, r)
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
	handler.ProductEdit(w, r)
}

func ProductUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handler.ProductUpdate(w, r)
	}
}
