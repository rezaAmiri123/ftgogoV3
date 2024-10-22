package queries

import (
	"context"
	
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
)

type GetOrder struct {
	ID string
}

type GetOrderHandler struct {
	orders domain.OrderRepository
}

func NewGetOrderHandler(orders domain.OrderRepository) GetOrderHandler {
	return GetOrderHandler{
		orders: orders,
	}
}

func (h GetOrderHandler) GetOrder(ctx context.Context, query GetOrder) (*domain.Order, error) {
	return h.orders.Find(ctx, query.ID)
}
