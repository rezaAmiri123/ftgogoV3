package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/order/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application/queries"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
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

func (a Application) CreateOrder(ctx context.Context, cmd commands.CreateOrder) (err error) {
	a.logger.Info().Msg("-->order.CreateOrder")
	defer func() { a.logger.Info().Err(err).Msg("<--order.CreateOrder") }()
	return a.App.CreateOrder(ctx, cmd)
}

func (a Application) GetOrder(ctx context.Context, query queries.GetOrder) (_ *domain.Order, err error) {
	a.logger.Info().Msg("-->order.GetOrder")
	defer func() { a.logger.Info().Err(err).Msg("<--order.GetOrder") }()
	return a.App.GetOrder(ctx, query)
}
