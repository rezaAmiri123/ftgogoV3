package domain

import (
	"context"
)

type CreateDelivery struct {
	DeliveryID   string
	RestaurantID string
	Address      Address
}

type DeliveryRepository interface {
	CreateDelivery(ctx context.Context, create CreateDelivery) error
}
