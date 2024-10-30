package commands

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/domain"
)

type UpdateRestaurantMenu struct {
	RestaurantID string
	MenuItems    []domain.MenuItem
}

type UpdateRestaurantMenuHandler struct {
	restaurants domain.RestaurantRepository
}

func NewUpdateRestaurantMenuHandler(restaurants domain.RestaurantRepository) UpdateRestaurantMenuHandler {
	return UpdateRestaurantMenuHandler{
		restaurants: restaurants,
	}
}

func (h UpdateRestaurantMenuHandler) UpdateRestaurantMenu(ctx context.Context, cmd UpdateRestaurantMenu) error {
	err := h.restaurants.UpdateMenu(ctx, domain.UpdateMenu{
		RestaurantID: cmd.RestaurantID,
		MenuItems:    cmd.MenuItems,
	})
	
	return err
}
