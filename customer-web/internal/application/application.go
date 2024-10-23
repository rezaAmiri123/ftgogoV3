package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		RegisterConsumer(ctx context.Context, cmd commands.RegisterConsumer) (string, error)
	}
	Queries interface{}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.RegisterConsumerHandler
	}
	appQueries struct{}
)

var _ App = (*Application)(nil)

func New(consumers domain.ConsumerRepository) *Application {
	return &Application{
		appCommands: appCommands{
			RegisterConsumerHandler: commands.NewRegisterConsumerHandler(consumers),
		},
		appQueries: appQueries{},
	}
}