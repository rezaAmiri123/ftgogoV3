package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
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
		ConfirmCreateTicket(ctx context.Context, cmd commands.ConfirmCreateTicket) error
		AcceptTicket(ctx context.Context, cmd commands.AcceptTicket) error
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
		commands.ConfirmCreateTicketHandler
		commands.AcceptTicketHandler
	}
	appQueries struct {
		queries.GetTicketHandler
	}
)

var _ App = (*Application)(nil)

func New(tickets domain.TicketRepository, domainPublisher ddd.EventPublisher) *Application {
	return &Application{
		appCommands: appCommands{
			CreateTicketHandler:        commands.NewCreateTicketHandler(tickets),
			ConfirmCreateTicketHandler: commands.NewConfirmCreateTicketHandler(tickets),
			AcceptTicketHandler:        commands.NewAcceptTicketHandler(tickets, domainPublisher),
		},
		appQueries: appQueries{
			GetTicketHandler: queries.NewGetTicketHandler(tickets),
		},
	}
}
