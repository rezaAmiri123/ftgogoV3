package domain

import (
	"context"
)

type RestaurantRepository interface {
	Find(ctx context.Context, restaurantID string) (*Restaurant, error)
}
