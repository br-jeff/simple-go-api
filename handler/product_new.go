package handler

import "net/http"

func ProductNew(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NewProduct", nil)

}
