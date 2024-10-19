package domain

import "context"

type RestaurantRepository interface {
	Save(ctx context.Context, restaurant *Restaurant) error
	Find(ctx context.Context, restaurantID string) (*Restaurant, error)
	Update(ctx context.Context, restaurant *Restaurant) error
}
