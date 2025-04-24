package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order/internal/order/repository/order"
)

func StartREST() {
	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		repo := order.OrderRepository{}

		orders, err := repo.FindAll()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)
	})

	fmt.Println("Starting rest server")
	http.ListenAndServe(":8080", nil)
}
