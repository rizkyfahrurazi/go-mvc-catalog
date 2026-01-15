package controllers

import (
	"net/http"

	"online-shop/config"
	"online-shop/helpers"
	"online-shop/models"
)

/* =========================
   CHECKOUT PROCESS
========================= */
func CheckoutProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/checkout", http.StatusSeeOther)
		return
	}

	// 🚫 wajib login
	userID, ok := helpers.GetUserID(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session, _ := helpers.GetSession(r)

	// 🛒 ambil cart dari session
	cart, total, err := models.GetCart(session.ID)
	if err != nil || len(cart) == 0 {
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}

	// 1️⃣ Create Order
	orderID, err := models.CreateOrder(config.DB, userID, total)
	if err != nil {
		http.Error(w, "Failed create order", http.StatusInternalServerError)
		return
	}

	// 2️⃣ Create Order Items
	for _, item := range cart {
		err := models.CreateOrderItem(
			config.DB,
			orderID,
			item.Product.ID,
			item.Product.Price,
			item.Qty,
		)
		if err != nil {
			http.Error(w, "Failed create order item", http.StatusInternalServerError)
			return
		}
	}

	// 3️⃣ Clear Cart
	models.ClearCart(session.ID)

	http.Redirect(w, r, "/checkout/success", http.StatusSeeOther)
}

/* =========================
   CHECKOUT SUCCESS
========================= */
func CheckoutSuccess(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Order Success",
		"Page":  "checkout-success",
	}

	helpers.Render(w, "checkout/success.html", data)
}
