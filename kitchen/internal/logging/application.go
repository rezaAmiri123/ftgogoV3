package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application/queries"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger zerolog.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}

func (a Application) CreateTicket(ctx context.Context, cmd commands.CreateTicket) (err error) {
	a.logger.Info().Msg("-->kitchen.CreateTicket")
	defer func() { a.logger.Info().Err(err).Msg("<--kitchen.CreateTicket") }()
	return a.App.CreateTicket(ctx, cmd)
}

func (a Application) GetTicket(ctx context.Context, query queries.GetTicket) (_ *domain.Ticket,err error){
	a.logger.Info().Msg("-->kitchen.GetTicket")
	defer func() { a.logger.Info().Err(err).Msg("<--kitchen.GetTicket") }()
	return a.App.GetTicket(ctx, query)
}

func (a Application) ConfirmCreateTicket(ctx context.Context, cmd commands.ConfirmCreateTicket) (err error) {
	a.logger.Info().Msg("-->kitchen.ConfirmCreateTicket")
	defer func() { a.logger.Info().Err(err).Msg("<--kitchen.ConfirmCreateTicket") }()
	return a.App.ConfirmCreateTicket(ctx, cmd)
}

func (a Application) AcceptTicket(ctx context.Context, cmd commands.AcceptTicket) (err error) {
	a.logger.Info().Msg("-->kitchen.AcceptTicket")
	defer func() { a.logger.Info().Err(err).Msg("<--kitchen.AcceptTicket") }()
	return a.App.AcceptTicket(ctx, cmd)
}
