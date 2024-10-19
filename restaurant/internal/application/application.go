package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/domain"
)

type (
	CreateRestaurant struct {
		ID      string
		Name    string
		Address domain.Address
	}
	GetRestaurant struct {
		ID string
	}
	UpdateMenuItem struct {
		RestaurantID string
		MenuItems    []domain.MenuItem
	}

	App interface {
		CreateRestaurant(ctx context.Context, create CreateRestaurant) error
		GetRestaurant(ctx context.Context, get GetRestaurant) (*domain.Restaurant, error)
		UpdateMenuItem(ctx context.Context, update UpdateMenuItem) error
	}

	Application struct {
		restaurants domain.RestaurantRepository
	}
)

var _ App = (*Application)(nil)

func New(restaurants domain.RestaurantRepository) *Application {
	return &Application{
		restaurants: restaurants,
	}
}

func (a Application) CreateRestaurant(ctx context.Context, create CreateRestaurant) error {
	Restaurant, err := domain.CreateRestaurant(create.ID, create.Name, create.Address)
	if err != nil {
		return err
	}

	return a.restaurants.Save(ctx, Restaurant)
}

func (a Application) GetRestaurant(ctx context.Context, get GetRestaurant) (*domain.Restaurant, error) {
	return a.restaurants.Find(ctx, get.ID)
}

func (a Application) UpdateMenuItem(ctx context.Context, update UpdateMenuItem) error {
	restaurant, err := a.restaurants.Find(ctx, update.RestaurantID)
	if err != nil {
		return err
	}
	err = restaurant.UpdateMenuItem(update.MenuItems)
	if err != nil {
		return err
	}

	return a.restaurants.Update(ctx, restaurant)
}
