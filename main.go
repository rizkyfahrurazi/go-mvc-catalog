package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"online-shop/config"
	"online-shop/models"
	"online-shop/routes"
)
func init() {
	gob.Register([]models.CartItem{})
	gob.Register(models.CartItem{})
}
func main() {
	config.ConnectDB()

	r := routes.SetupRoutes()

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
