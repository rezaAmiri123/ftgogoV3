package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application/queries"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateTicket(ctx context.Context, cmd commands.CreateTicket) error
	}
	Queries interface {
		GetTicket(ctx context.Context, query queries.GetTicket) (*domain.Ticket, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateTicketHandler
	}
	appQueries struct {
		queries.GetTicketHandler
	}
)

var _ App = (*Application)(nil)

func New(tickets domain.TicketRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateTicketHandler: commands.NewCreateTicketHandler(tickets),
		},
		appQueries: appQueries{
			GetTicketHandler: queries.NewGetTicketHandler(tickets),
		},
	}
}
