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
)

type OrderRepository interface {
	Create(ctx context.Context, create CreateOrder) (string, error)
}
