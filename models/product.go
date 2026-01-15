package models

import (
	"online-shop/config"
)

type Product struct {
	ID    int
	Name  string
	Price int
	Image string
}

func GetAllProducts() ([]Product, error) {
	rows, err := config.DB.Query(`
		SELECT id, name, price, image 
		FROM products
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Image)
		products = append(products, p)
	}

	return products, nil
}

func GetProductByID(id int) (*Product, error) {
	row := config.DB.QueryRow(`
		SELECT id, name, price, image 
		FROM products 
		WHERE id = ?
	`, id)

	var p Product
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Image)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
