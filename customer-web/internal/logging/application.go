package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application/commands"
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
	a.logger.Info().Msg("-->consumer.RegisterConsumer")
	defer func() { a.logger.Info().Err(err).Msg("<--consumer.RegisterConsumer") }()
	return a.App.RegisterConsumer(ctx, cmd)
}
