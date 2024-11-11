package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rs/zerolog"
)

type ConsumerHandlers[T ddd.Event] struct {
	logger zerolog.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*ConsumerHandlers[ddd.Event])(nil)

func NewConsumerHandlers(logger zerolog.Logger) ConsumerHandlers[ddd.Event] {
	return ConsumerHandlers[ddd.Event]{
		logger: logger,
	}
}

func (h ConsumerHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case consumerpb.ConsumerRegisteredEvent:
		return h.onConsumerRegistered(ctx, event)
	}
	return nil
}

func (h ConsumerHandlers[T]) onConsumerRegistered(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*consumerpb.ConsumerRegistred)
	h.logger.Debug().Msgf(`ID: %s, Name: %s`, payload.GetId(), payload.GetName())
	
	// return h.app.RegisterAccount(ctx,RegisterAccount{
	// 	ID: payload.GetId(),
	// 	Name: payload.GetName(),
	// })
	return nil
}
