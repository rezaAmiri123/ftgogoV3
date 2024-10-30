package commands

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/domain"
)

type CreateRestaurant struct {
	Name      string
	Address   domain.Address
}

type CreateRestaurantHandler struct {
	restaurants domain.RestaurantRepository
}

func NewCreateRestaurantHandler(restaurants domain.RestaurantRepository) CreateRestaurantHandler {
	return CreateRestaurantHandler{
		restaurants: restaurants,
	}
}

func (h CreateRestaurantHandler) CreateRestaurant(ctx context.Context, cmd CreateRestaurant) (string, error) {
	restaurantID, err := h.restaurants.Create(ctx, domain.CreateRestaurant{
		Name: cmd.Name,
		Address: cmd.Address,
	})
	if err != nil {
		return "", err
	}

	return restaurantID, nil
}
