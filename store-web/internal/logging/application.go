package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/application/commands"
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

func (a Application) CreateRestaurant(ctx context.Context, cmd commands.CreateRestaurant) (_ string, err error) {
	a.logger.Info().Msg("-->store-web.CreateRestaurant")
	defer func() { a.logger.Info().Err(err).Msg("<--store-web.CreateRestaurant") }()
	return a.App.CreateRestaurant(ctx, cmd)
}

