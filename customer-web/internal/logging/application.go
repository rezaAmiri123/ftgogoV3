package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application/queries"
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

func (a Application) RegisterConsumer(ctx context.Context, cmd commands.RegisterConsumer) (_ string, err error) {
	a.logger.Info().Msg("-->customer-web.RegisterConsumer")
	defer func() { a.logger.Info().Err(err).Msg("<--customer-web.RegisterConsumer") }()
	return a.App.RegisterConsumer(ctx, cmd)
}

// (ctx context.Context,  ) (*domain.Consumer, error)
func (a Application) GetConsumer(ctx context.Context,query queries.GetConsumer) (_ *domain.Consumer, err error) {
	a.logger.Info().Msg("-->customer-web.GetConsumer")
	defer func() { a.logger.Info().Err(err).Msg("<--customer-web.GetConsumer") }()
	return a.App.GetConsumer(ctx, query)
}
