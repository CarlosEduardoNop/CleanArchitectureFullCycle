package order

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"order/internal/order/entity"
)

type OrderRepository struct {
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=root password=root dbname=order_db sslmode=disable")

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
