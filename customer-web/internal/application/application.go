package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application/queries"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		RegisterConsumer(ctx context.Context, cmd commands.RegisterConsumer) (string, error)
		AddConsumerAddress(ctx context.Context, cmd commands.AddConsumerAddress) error
		CreateOrder(ctx context.Context, cmd commands.CreateOrder) (string, error)
	}
	Queries interface {
		GetConsumer(ctx context.Context, query queries.GetConsumer) (*domain.Consumer, error)
		GetOrder(ctx context.Context, query queries.GetOrder) (*domain.Order, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.RegisterConsumerHandler
		commands.AddConsumerAddressHandler
		commands.CreateOrderHandler
	}
	appQueries struct {
		queries.GetConsumerHandler
		queries.GetOrderHandler
	}
)

var _ App = (*Application)(nil)

func New(consumers domain.ConsumerRepository, orders domain.OrderRepository) *Application {
	return &Application{
		appCommands: appCommands{
			RegisterConsumerHandler:   commands.NewRegisterConsumerHandler(consumers),
			AddConsumerAddressHandler: commands.NewAddConsumerAddressHandler(consumers),
			CreateOrderHandler:        commands.NewCreateOrderHandler(consumers, orders),
		},
		appQueries: appQueries{
			GetConsumerHandler: queries.NewGetConsumerHandler(consumers),
			GetOrderHandler:    queries.NewGetOrderHandler(orders),
		},
	}
}
