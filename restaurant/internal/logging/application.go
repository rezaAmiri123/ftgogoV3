package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/domain"
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




func (a Application)CreateRestaurant(ctx context.Context, create application.CreateRestaurant) (err error) {
	a.logger.Info().Msg("-->restaurant.CreateRestaurant")
	defer func() { a.logger.Info().Err(err).Msg("<--restaurant.CreateRestaurant") }()
	return a.App.CreateRestaurant(ctx, create)
}

func (a Application)GetRestaurant(ctx context.Context, get application.GetRestaurant) (_ *domain.Restaurant, err error){
	a.logger.Info().Msg("-->restaurant.GetRestaurant")
	defer func() { a.logger.Info().Err(err).Msg("<--restaurant.GetRestaurant") }()
	return a.App.GetRestaurant(ctx, get)
}

func (a Application)UpdateMenuItem(ctx context.Context, update application.UpdateMenuItem)(err error) {
	a.logger.Info().Msg("-->restaurant.UpdateMenuItem")
	defer func() { a.logger.Info().Err(err).Msg("<--restaurant.UpdateMenuItem") }()
	return a.App.UpdateMenuItem(ctx, update)
}
