package routes

import (
	"net/http"
	"online-shop/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// pages
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/catalog", controllers.Catalog).Methods("GET")
	r.HandleFunc("/product/{id}", controllers.ProductDetail).Methods("GET")

	// auth
	r.HandleFunc("/login", controllers.Login).Methods("GET", "POST")
	r.HandleFunc("/register", controllers.Register).Methods("GET", "POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")

	// CART
	r.HandleFunc("/cart", controllers.CartPage).Methods("GET")
	r.HandleFunc("/add-to-cart", controllers.AddToCart).Methods("GET")
	r.HandleFunc("/cart/increase", controllers.IncreaseQty).Methods("GET")
	r.HandleFunc("/cart/decrease", controllers.DecreaseQty).Methods("GET")
	r.HandleFunc("/checkout", controllers.CheckoutProcess).Methods("POST")
r.HandleFunc("/checkout/process", controllers.CheckoutProcess)

	r.HandleFunc("/checkout/success", controllers.CheckoutSuccess).Methods("GET")



	// static files
	r.PathPrefix("/public/").
		Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	return r
}
