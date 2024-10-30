package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateRestaurant(ctx context.Context, cmd commands.CreateRestaurant) (string, error)
		UpdateRestaurantMenu(ctx context.Context, cmd commands.UpdateRestaurantMenu) error
	}
	Queries interface {
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateRestaurantHandler
		commands.UpdateRestaurantMenuHandler
	}
	appQueries struct {
	}
)

var _ App = (*Application)(nil)

func New(restaurants domain.RestaurantRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateRestaurantHandler:     commands.NewCreateRestaurantHandler(restaurants),
			UpdateRestaurantMenuHandler: commands.NewUpdateRestaurantMenuHandler(restaurants),
		},
		appQueries: appQueries{},
	}
}
