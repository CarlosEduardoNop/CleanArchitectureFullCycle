package usecase

import (
	"order/app/internal/order/entity"
	"order/app/internal/order/repository/order"
)

type CreateOrderService struct{}

func (s *CreateOrderService) Create(o *entity.Order) error {
	repository := order.OrderRepository{}
	return repository.Create(o)
}
