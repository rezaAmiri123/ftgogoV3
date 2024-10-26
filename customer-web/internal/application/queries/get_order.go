package queries

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
	"github.com/stackus/errors"
)

type GetOrder struct {
	OrderID    string
	ConsumerID string
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
	order, err := h.orders.Find(ctx, domain.FindOrder{
		OrderID: query.OrderID,
	})
	if err != nil {
		return nil, err
	}
	if order.ConsumerID != query.ConsumerID{
		// being opaque intentionally; Could also send a permission denied error
		return nil, errors.Wrap(errors.ErrNotFound, "order not found")
	}

	return order, nil
}
