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

func (a Application) SetCourierAvailability(ctx context.Context, cmd commands.SetCourierAvailability) (err error) {
	a.logger.Info().Msg("-->delivery.SetCourierAvailability")
	defer func() { a.logger.Info().Err(err).Msg("<--delivery.SetCourierAvailability") }()
	return a.App.SetCourierAvailability(ctx, cmd)
}

func (a Application) ScheduleDelivery(ctx context.Context, cmd commands.ScheduleDelivery) (err error) {
	a.logger.Info().Msg("-->delivery.ScheduleDelivery")
	defer func() { a.logger.Info().Err(err).Msg("<--delivery.ScheduleDelivery") }()
	return a.App.ScheduleDelivery(ctx, cmd)
}
