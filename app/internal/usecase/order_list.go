package usecase

import (
	"order/app/internal/order/entity"
	"order/app/internal/order/repository/order"
)

type GetOrderService struct{}

func (s *GetOrderService) ListOrders() ([]entity.Order, error) {
	repository := order.OrderRepository{}
	return repository.FindAll()
}
