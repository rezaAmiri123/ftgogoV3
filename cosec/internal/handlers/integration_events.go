package handlers

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/cosec/internal/models"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/errorsotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/sec"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type integrationHandlers[T ddd.Event] struct {
	orchestrator sec.Orchestrator[*models.CreateOrderData]
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(reg registry.Registry, orchestrator sec.Orchestrator[*models.CreateOrderData], mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewEventHandler(reg, integrationHandlers[ddd.Event]{
		orchestrator: orchestrator,
	}, mws...)
}

func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) (err error) {
	evtMsgHandler := am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) error {
		return handlers.HandleMessage(ctx, msg)
	})

	_, err = subscriber.Subscribe(orderpb.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderpb.OrderCreatedEvent,
	}, am.GroupName("cosec-orders"))
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
	case orderpb.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	}

	return nil
}

func (h integrationHandlers[T]) onOrderCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*orderpb.OrderCreated)

	var total int
	items := make([]models.LineItem, len(payload.GetItems()))
	for i, item := range payload.GetItems() {
		items[i] = models.LineItem{
			MenuItemID: item.GetMenuItemId(),
			Name:       item.GetName(),
			Price:      int(item.GetPrice()),
			Quantity:   int(item.GetQuantity()),
		}
		total += int(item.GetPrice()) * int(item.GetQuantity())
	}
	data := models.CreateOrderData{
		OrderID:      payload.GetOrderID(),
		ConsumerID:   payload.GetConsumerID(),
		RestaurantID: payload.GetRestaurantID(),
		LineItems:    items,
		OrderTotal:   total,
	}

	// Start the CreateOrderSaga
	return h.orchestrator.Start(ctx, event.ID(), &data)
}
