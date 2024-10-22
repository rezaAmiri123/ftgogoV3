package commands

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
)

type CreateOrder struct {
	ID           string
	ConsumerID   string
	RestaurantID string
	DeliverAt    time.Time
	DeliverTo    domain.Address
	LineItems    map[string]int
}

type CreateOrderHandler struct {
	orders      domain.OrderRepository
	restaurants domain.RestaurantRepository
}

func NewCreateOrderHandler(
	orders domain.OrderRepository,
	restaurants domain.RestaurantRepository,
) CreateOrderHandler {
	return CreateOrderHandler{
		orders:      orders,
		restaurants: restaurants,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	restaurant, err := h.restaurants.Find(ctx, cmd.RestaurantID)
	if err != nil {
		return err
	}
	lineItems, err := restaurant.LineItems(cmd.LineItems)
	if err != nil {
		return err
	}
	order, err := domain.CreateOrder(cmd.ID, cmd.ConsumerID, cmd.RestaurantID, lineItems, cmd.DeliverAt, cmd.DeliverTo)
	if err != nil {
		return err
	}
	return h.orders.Save(ctx, order)
}
