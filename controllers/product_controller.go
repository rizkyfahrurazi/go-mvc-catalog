package controllers

import (
	"net/http"
	"strconv"

	"online-shop/helpers"
	"online-shop/models"

	"github.com/gorilla/mux"
)

func ProductDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	product, _ := models.GetProductByID(id)
	if product == nil {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"Title":      product.Name + " - Jianaka",
		"IsLoggedIn": helpers.IsLoggedIn(r),
		"Page":       "product",
		"Product":    product,
	}

	helpers.Render(w, "product/detail.html", data)
}
