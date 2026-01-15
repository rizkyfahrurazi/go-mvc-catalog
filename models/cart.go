package models

import (
	"online-shop/config"
)

type CartItem struct {
	ID       int
	Product  Product
	Qty      int
	Subtotal int
}

/* =========================
   ADD / UPDATE CART
========================= */
func AddToCart(sessionID string, productID int) error {
	_, err := config.DB.Exec(`
		INSERT INTO cart_items (session_id, product_id, qty)
		VALUES (?, ?, 1)
		ON DUPLICATE KEY UPDATE qty = qty + 1
	`, sessionID, productID)

	return err
}

/* =========================
   GET CART ITEMS
========================= */
func GetCart(sessionID string) ([]CartItem, int, error) {
	rows, err := config.DB.Query(`
		SELECT 
			c.id, c.qty,
			p.id, p.name, p.price, p.image
		FROM cart_items c
		JOIN products p ON p.id = c.product_id
		WHERE c.session_id = ?
	`, sessionID)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var cart []CartItem
	total := 0

	for rows.Next() {
		var item CartItem
		var product Product

		rows.Scan(
			&item.ID,
			&item.Qty,
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Image,
		)

		item.Product = product
		item.Subtotal = product.Price * item.Qty
		total += item.Subtotal

		cart = append(cart, item)
	}

	return cart, total, nil
}

/* =========================
   INCREASE / DECREASE
========================= */
func IncreaseQty(sessionID string, productID int) {
	config.DB.Exec(`
		UPDATE cart_items 
		SET qty = qty + 1
		WHERE session_id = ? AND product_id = ?
	`, sessionID, productID)
}

func DecreaseQty(sessionID string, productID int) {
	config.DB.Exec(`
		UPDATE cart_items 
		SET qty = qty - 1
		WHERE session_id = ? AND product_id = ? AND qty > 1
	`, sessionID, productID)
}

/* =========================
   REMOVE ITEM
========================= */
func RemoveItem(sessionID string, productID int) {
	config.DB.Exec(`
		DELETE FROM cart_items 
		WHERE session_id = ? AND product_id = ?
	`, sessionID, productID)
}

func ClearCart(sessionID string) {
	config.DB.Exec(`
		DELETE FROM cart_items WHERE session_id = ?
	`, sessionID)
}

