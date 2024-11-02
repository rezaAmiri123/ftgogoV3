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
		AcceptTicket(ctx context.Context, cmd commands.AcceptTicket) error
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
		commands.AcceptTicketHandler
	}
	appQueries struct {
	}
)

var _ App = (*Application)(nil)

func New(restaurants domain.RestaurantRepository, kitchens domain.KitchenRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateRestaurantHandler:     commands.NewCreateRestaurantHandler(restaurants),
			UpdateRestaurantMenuHandler: commands.NewUpdateRestaurantMenuHandler(restaurants),
			AcceptTicketHandler:         commands.NewAcceptTicketHandler(kitchens),
		},
		appQueries: appQueries{},
	}
}
