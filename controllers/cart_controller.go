package controllers

import (
	"net/http"
	"strconv"

	"online-shop/helpers"
	"online-shop/models"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	session, _ := helpers.GetSession(r)

	models.AddToCart(session.ID, id)

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func CartPage(w http.ResponseWriter, r *http.Request) {
	session, _ := helpers.GetSession(r)

	cart, total, _ := models.GetCart(session.ID)

	data := map[string]interface{}{
		"Title":	"Cart",
		"IsLoggedIn":	helpers.IsLoggedIn(r),
		"Cart":  cart,
		"Total": total,
	}

	helpers.Render(w, "cart/index.html", data)
}

func IncreaseQty(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	session, _ := helpers.GetSession(r)

	models.IncreaseQty(session.ID, id)
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func DecreaseQty(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	session, _ := helpers.GetSession(r)

	models.DecreaseQty(session.ID, id)
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}
