package controllers

import (
	"net/http"
	"online-shop/helpers"
	"online-shop/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	products, _ := models.GetAllProducts()

	data := map[string]interface{}{
		"Title":      "Jianaka",
		"IsLoggedIn": helpers.IsLoggedIn(r),
		"Page":       "home",
		"Products":   products,
	}

	helpers.Render(w, "home/index.html", data)
}
