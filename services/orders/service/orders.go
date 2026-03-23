package service

import (
	"Order/services/common/genproto/orders"
	"context"
)

type OrderService struct {
}

var ordersData = []*orders.Order{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersData = append(ordersData, order)

	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersData
}
