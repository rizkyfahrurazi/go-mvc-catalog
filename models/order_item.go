package models

import "database/sql"

func CreateOrderItem(
	db *sql.DB,
	orderID int,
	productID int,
	price int,
	qty int,
) error {

	subtotal := price * qty

	_, err := db.Exec(`
		INSERT INTO order_items 
		(order_id, product_id, price, qty, subtotal)
		VALUES (?, ?, ?, ?, ?)`,
		orderID, productID, price, qty, subtotal,
	)

	return err
}
