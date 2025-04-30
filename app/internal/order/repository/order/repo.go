package order

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"order/app/internal/order/entity"
)

type OrderRepository struct {
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {
	db, err := sql.Open("pgx", "host=postgres port=5432 user=root password=root dbname=order_db sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, item, price FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var o entity.Order
		err := rows.Scan(&o.ID, &o.Item, &o.Price)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (r *OrderRepository) Create(o *entity.Order) error {
	db, err := sql.Open("pgx", "host=postgres port=5432 user=root password=root dbname=order_db sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO orders (item, price) VALUES ($1, $2) RETURNING id"

	err = db.QueryRow(query, o.Item, o.Price).Scan(&o.ID)

	return err
}
