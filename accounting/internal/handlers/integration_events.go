package handlers

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type integrationHandlers[T ddd.Event] struct {
	app application.App
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationHandlers(app application.App) ddd.EventHandler[ddd.Event] {
	return integrationHandlers[ddd.Event]{
		app: app,
	}
}

func RegisterIntegrationEventHandlers(subscriber am.RawMessageStream, handlers am.RawMessageHandler) (err error) {
	evtMsgHandler := am.RawMessageHandlerFunc(func(ctx context.Context, msg am.IncomingRawMessage) error {
		return handlers.HandleMessage(ctx, msg)
	})

	_, err = subscriber.Subscribe(consumerpb.ConsumerAggregateChannel, evtMsgHandler, am.MessageFilter{
		consumerpb.ConsumerRegisteredEvent,
	}, am.GroupName("accounting-consumer"))
	return err
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case consumerpb.ConsumerRegisteredEvent:
		return h.onConsumerRegistered(ctx, event)
	}

	return nil
}

func (h integrationHandlers[T]) onConsumerRegistered(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*consumerpb.ConsumerRegistred)
	return h.app.RegisterAccount(ctx, application.RegisterAccount{
		ID:   payload.GetId(),
		Name: payload.GetName(),
	})
}
