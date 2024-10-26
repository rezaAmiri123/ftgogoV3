package commands

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
)

type CreateOrder struct {
	ConsumerID   string
	RestaurantID string
	AddressID    string
	LineItems    domain.MenuItemQuantities
}

type CreateOrderHandler struct {
	consumers domain.ConsumerRepository
	orders    domain.OrderRepository
}

func NewCreateOrderHandler(consumers domain.ConsumerRepository, orders domain.OrderRepository) CreateOrderHandler {
	return CreateOrderHandler{
		consumers: consumers,
		orders:    orders,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) (string, error) {
	address, err := h.consumers.FindAddress(ctx, domain.FindConsumerAddress{
		ConsumerID: cmd.ConsumerID,
		AddressID:  cmd.AddressID,
	})
	if err != nil {
		return "", err
	}

	return h.orders.Create(ctx, domain.CreateOrder{
		ConsumerID:   cmd.ConsumerID,
		RestaurantID: cmd.RestaurantID,
		DeliverAt:    time.Now().Add(30 * time.Minute),
		DeliverTo:    *address,
		LineItems:    cmd.LineItems,
	})
}
