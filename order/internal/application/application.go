package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/order/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application/queries"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
)
type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateOrder(ctx context.Context, cmd commands.CreateOrder) error
	}
	Queries interface {
		GetOrder(ctx context.Context, query queries.GetOrder) (*domain.Order, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateOrderHandler
	}
	appQueries struct {
		queries.GetOrderHandler
	}
)

var _ App = (*Application)(nil)

func New(orders domain.OrderRepository, restaurants domain.RestaurantRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateOrderHandler: commands.NewCreateOrderHandler(orders,restaurants),
		},
		appQueries: appQueries{
			GetOrderHandler: queries.NewGetOrderHandler(orders),
		},
	}
}