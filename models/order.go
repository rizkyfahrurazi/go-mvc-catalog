package models

import "database/sql"

type Order struct {
	ID     int
	UserID int
	Total  int
	Status string
}

func CreateOrder(db *sql.DB, userID int, total int) (int, error) {
	result, err := db.Exec(
		"INSERT INTO orders (user_id, total) VALUES (?, ?)",
		userID, total,
	)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return int(id), nil
}
