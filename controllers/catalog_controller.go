package controllers

import (
	"net/http"
	"online-shop/helpers"
	"online-shop/models"
)

func Catalog(w http.ResponseWriter, r *http.Request) {
	products, _ := models.GetAllProducts()

	data := map[string]interface{}{
		"Title":      "Catalog - Jianaka",
		"IsLoggedIn": helpers.IsLoggedIn(r),
		"Page":       "catalog",
		"Products":   products,
	}

	helpers.Render(w, "catalog/index.html", data)
}
