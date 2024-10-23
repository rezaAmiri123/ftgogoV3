package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
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

func (a Application) RegisterConsumer(ctx context.Context, register application.RegisterConsumer) (err error) {
	a.logger.Info().Msg("-->consumer.RegisterConsumer")
	defer func() { a.logger.Info().Err(err).Msg("<--consumer.RegisterConsumer") }()
	return a.App.RegisterConsumer(ctx, register)
}

func (a Application)GetConsumer(ctx context.Context, get application.GetConsumer) (_ *domain.Consumer,err error) {
	a.logger.Info().Msg("-->consumer.GetConsumer")
	defer func() { a.logger.Info().Err(err).Msg("<--consumer.GetConsumer") }()
	return a.App.GetConsumer(ctx, get)
}

func (a Application)UpdateConsumerAddress(ctx context.Context, update application.UpdateConsumerAddress) (err error){
	a.logger.Info().Msg("-->consumer.UpdateConsumerAddress")
	defer func() { a.logger.Info().Err(err).Msg("<--consumer.UpdateConsumerAddress") }()
	return a.App.UpdateConsumerAddress(ctx, update)
}

func (a Application)RemoveConsumerAddress(ctx context.Context, remove application.RemoveConsumerAddress) (err error){
	a.logger.Info().Msg("-->consumer.RemoveConsumerAddress")
	defer func() { a.logger.Info().Err(err).Msg("<--consumer.RemoveConsumerAddress") }()
	return a.App.RemoveConsumerAddress(ctx, remove)
}

func (a Application)GetConsumerAddress(ctx context.Context, get application.GetConsumerAddress) (_ domain.Address,err error){
	a.logger.Info().Msg("-->consumer.GetConsumerAddress")
	defer func() { a.logger.Info().Err(err).Msg("<--consumer.GetConsumerAddress") }()
	return a.App.GetConsumerAddress(ctx, get)
}

func (a Application)ValidateOrderByConsumer(ctx context.Context, validate application.ValidateOrderByConsumer) (err error){
	a.logger.Info().Msg("-->consumer.ValidateOrderByConsumer")
	defer func() { a.logger.Info().Err(err).Msg("<--consumer.ValidateOrderByConsumer") }()
	return a.App.ValidateOrderByConsumer(ctx, validate)
}
