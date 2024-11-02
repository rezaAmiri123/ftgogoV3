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
		CreateDelivery(ctx context.Context, cmd commands.CreateDelivery) error
		SetCourierAvailability(ctx context.Context, cmd commands.SetCourierAvailability) error
		ScheduleDelivery(ctx context.Context, cmd commands.ScheduleDelivery) error
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
		commands.SetCourierAvailabilityHandler
		commands.ScheduleDeliveryHandler
	}
	appQueries struct {
		queries.GetDeliveryHandler
	}
)

var _ App = (*Application)(nil)

func New(
	deliveries domain.DeliveryRepository,
	couriers domain.CourierRepository,
	restaurants domain.RestaurantRepository,
) *Application {
	return &Application{
		appCommands: appCommands{
			CreateDeliveryHandler:         commands.NewCreateDeliveryHandler(deliveries, restaurants),
			SetCourierAvailabilityHandler: commands.NewSetCourierAvailabilityHandler(couriers),
			ScheduleDeliveryHandler:       commands.NewScheduleDeliveryHandler(deliveries, couriers),
		},
		appQueries: appQueries{
			GetDeliveryHandler: queries.NewGetDeliveryHandler(deliveries),
		},
	}
}
