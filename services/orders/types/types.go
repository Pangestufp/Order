package types

import (
	"Order/services/common/genproto/orders"
	"context"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
