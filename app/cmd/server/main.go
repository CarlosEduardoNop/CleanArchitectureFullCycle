package main

import (
	_ "github.com/lib/pq"
	"order/app/internal/order/delivery/graphql"
	"order/app/internal/order/delivery/grpcserver"
	"order/app/internal/order/delivery/rest"
)

func main() {
	go grpcserver.StartGRPC()
	go graphql.StartGraphQL()
	//
	rest.StartREST()
}
