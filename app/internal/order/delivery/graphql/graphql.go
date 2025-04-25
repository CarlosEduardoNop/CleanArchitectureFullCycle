package graphql

import (
	"encoding/json"
	"fmt"
	gql "github.com/graphql-go/graphql"
	"net/http"
	"order/app/internal/order/repository/order"
)

func StartGraphQL() {
	orderType := gql.NewObject(gql.ObjectConfig{
		Name: "Order",
		Fields: gql.Fields{
			"id":    &gql.Field{Type: gql.Int},
			"item":  &gql.Field{Type: gql.String},
			"price": &gql.Field{Type: gql.Float},
		},
	})

	repo := order.OrderRepository{}

	rootQuery := gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"listOrders": &gql.Field{
				Type: gql.NewList(orderType),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return repo.FindAll()
				},
			},
		},
	})

	schema, _ := gql.NewSchema(gql.SchemaConfig{
		Query: rootQuery,
	})

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query string `json:"query"`
		}

		json.NewDecoder(r.Body).Decode(&params)

		result := gql.Do(gql.Params{
			Schema:        schema,
			RequestString: params.Query,
		})
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Starting graphql")
	http.ListenAndServe(":8081", nil)
}
