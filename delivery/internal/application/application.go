package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application/queries"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
)


type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateDelivery(ctx context.Context, cmd commands.CreateDelivery)error
	}
	Queries interface {
		GetDelivery(ctx context.Context, query queries.GetDelivery) (*domain.Delivery, error) 
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateDeliveryHandler
	}
	appQueries struct {
		queries.GetDeliveryHandler
	}
)

var _ App = (*Application)(nil)

func New(deliveries domain.DeliveryRepository, restaurants domain.RestaurantRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateDeliveryHandler: commands.NewCreateDeliveryHandler(deliveries,restaurants),
		},
		appQueries: appQueries{
			GetDeliveryHandler: queries.NewGetDeliveryHandler(deliveries),
		},
	}
}
