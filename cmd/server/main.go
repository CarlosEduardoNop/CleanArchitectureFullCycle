package main

import (
	"order/internal/order/delivery/graphql"
	"order/internal/order/delivery/grpcserver"
	"order/internal/order/delivery/rest"

	_ "github.com/lib/pq"
)

func main() {
	go grpcserver.StartGRPC()
	go graphql.StartGraphQL()
	//
	rest.StartREST()
}
