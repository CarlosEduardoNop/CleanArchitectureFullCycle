package migrations

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateOrderTable() {
	db, err := sql.Open("pgx", "host=postgres port=5432 user=root password=root dbname=order_db sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS orders (id SERIAL PRIMARY KEY,item VARCHAR(255),price float);")
	if err != nil {
		log.Fatalf("Erro ao executar SQL: %v", err)
	}

	fmt.Println("Migration executada com sucesso!")
}
