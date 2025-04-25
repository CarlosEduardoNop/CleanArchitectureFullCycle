package grpcserver

import (
	"google.golang.org/grpc"
	"log"
	"net"
	order2 "order/app/internal/order/delivery/grpcserver/order"
)

func StartGRPC() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	o := order2.Server{}

	grpcServer := grpc.NewServer()

	order2.RegisterOrderServiceServer(grpcServer, &o)
	log.Println("gRPC server running on port 50051")
	grpcServer.Serve(lis)
}
