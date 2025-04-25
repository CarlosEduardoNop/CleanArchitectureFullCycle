package order

import (
	"context"
	"order/app/internal/order/repository/order"
)

type Server struct {
	UnimplementedOrderServiceServer
}

func (s *Server) ListOrders(ctx context.Context, req *ListOrdersRequest) (*ListOrdersResponse, error) {
	repo := order.OrderRepository{}

	orders, err := repo.FindAll()
	if err != nil {
		return nil, err
	}

	var pbOrders []*Order
	for _, o := range orders {
		pbOrders = append(pbOrders, &Order{
			Id:    o.ID,
			Item:  o.Item,
			Price: o.Price,
		})
	}

	return &ListOrdersResponse{Orders: pbOrders}, nil
}
