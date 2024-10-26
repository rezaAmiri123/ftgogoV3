package domain

import (
	"context"
	"time"
)

type (
	CreateOrder struct {
		ConsumerID   string
		RestaurantID string
		DeliverAt    time.Time
		DeliverTo    Address
		LineItems    MenuItemQuantities
	}
	FindOrder struct {
		OrderID string
	}
)

type OrderRepository interface {
	Create(ctx context.Context, create CreateOrder) (string, error)
	Find(ctx context.Context, find FindOrder) (*Order, error)
}
