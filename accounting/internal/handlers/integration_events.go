package handlers

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/errorsotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type integrationHandlers[T ddd.Event] struct {
	app application.App
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationHandlers(reg registry.Registry, app application.App, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewEventHandler(reg, integrationHandlers[ddd.Event]{
		app: app,
	}, mws...)
}

func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) (err error) {
	evtMsgHandler := am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) error {
		return handlers.HandleMessage(ctx, msg)
	})

	_, err = subscriber.Subscribe(consumerpb.ConsumerAggregateChannel, evtMsgHandler, am.MessageFilter{
		consumerpb.ConsumerRegisteredEvent,
	}, am.GroupName("accounting-consumer"))
	return err
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling integration event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled integration event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling integration event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))

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
