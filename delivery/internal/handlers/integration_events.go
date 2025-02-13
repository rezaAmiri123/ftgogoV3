package handlers

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/errorsotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
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

	_, err = subscriber.Subscribe(kitchenpb.TicketAggregateChannel, evtMsgHandler, am.MessageFilter{
		kitchenpb.TicketAcceptedEvent,
	}, am.GroupName("delivery-tickets"))
	if err != nil {
		return err
	}
	_, err = subscriber.Subscribe(orderpb.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderpb.OrderCreatedEvent,
	}, am.GroupName("delivery-orders"))
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
	case kitchenpb.TicketAcceptedEvent:
		return h.onTicketAccepted(ctx, event)
	case orderpb.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	}
	return nil
}

func (h integrationHandlers[T]) onTicketAccepted(ctx context.Context, event T) error {
	payload := event.Payload().(*kitchenpb.TicketAccepted)
	return h.app.ScheduleDelivery(ctx, commands.ScheduleDelivery{
		ID:      payload.GetOrderID(),
		ReadyBy: payload.ReadyBy.AsTime(),
	})
}

func (h integrationHandlers[T]) onOrderCreated(ctx context.Context, event T) error {
	payload := event.Payload().(*orderpb.OrderCreated)
	return h.app.CreateDelivery(ctx, commands.CreateDelivery{
		ID:              payload.GetOrderID(),
		RestaurantID:    payload.GetRestaurantID(),
		DeliveryAddress: h.toAddressDomain(payload.GetAddress()),
	})
}

func (integrationHandlers[T]) toAddressDomain(address *orderpb.Address) domain.Address {
	if address == nil {
		return domain.Address{}
	}
	return domain.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}
