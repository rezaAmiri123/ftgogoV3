package domain

import "context"

type (
	CreateRestaurant struct {
		Name      string
		Address   Address
	}
	UpdateMenu struct{
		RestaurantID string
		MenuItems []MenuItem
	}
)

type RestaurantRepository interface {
	Create(ctx context.Context,create CreateRestaurant) (string, error)
	UpdateMenu(ctx context.Context,update UpdateMenu)error
}