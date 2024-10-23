package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application/queries"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
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

func (a Application) CreateDelivery(ctx context.Context, cmd commands.CreateDelivery) (err error) {
	a.logger.Info().Msg("-->delivery.CreateDelivery")
	defer func() { a.logger.Info().Err(err).Msg("<--delivery.CreateDelivery") }()
	return a.App.CreateDelivery(ctx, cmd)
}

func (a Application) GetDelivery(ctx context.Context, query queries.GetDelivery) (_ *domain.Delivery,err error){
	a.logger.Info().Msg("-->delivery.GetDelivery")
	defer func() { a.logger.Info().Err(err).Msg("<--delivery.GetDelivery") }()
	return a.App.GetDelivery(ctx, query)
}
