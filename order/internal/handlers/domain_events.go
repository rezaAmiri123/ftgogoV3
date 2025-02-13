package handlers

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/errorsotel"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type domainHandlers[T ddd.Event] struct {
	publisher am.EventPublisher
}

var _ ddd.EventHandler[ddd.Event] = (*domainHandlers[ddd.Event])(nil)

func NewDomainEventHandlers(publisher am.EventPublisher) ddd.EventHandler[ddd.Event] {
	return &domainHandlers[ddd.Event]{
		publisher: publisher,
	}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handlers ddd.EventHandler[ddd.Event]) {
	subscriber.Subscribe(handlers,
		domain.OrderCreatedEvent,
	)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time){
		if err!= nil{
			span.AddEvent(
				"Encountered an error handling model event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handling domain event", trace.WithAttributes(
			attribute.Int64("TookMS", int64(time.Since(started).Milliseconds())),
		))
	}(time.Now())
	
	span.AddEvent("Handling domain event",trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))
	
	switch event.EventName() {
	case domain.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	}
	return nil
}

func (h domainHandlers[T]) onOrderCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.Order)
	items := make([]*orderpb.OrderCreated_Item, len(payload.LineItems))
	for i, item := range payload.LineItems {
		items[i] = &orderpb.OrderCreated_Item{
			MenuItemId: item.MenuItemID,
			Name:       item.Name,
			Price:      int64(item.Price),
			Quantity:   int64(item.Quantity),
		}
	}

	return h.publisher.Publish(ctx, orderpb.OrderAggregateChannel,
		ddd.NewEvent(orderpb.OrderCreatedEvent, &orderpb.OrderCreated{
			OrderID:      payload.ID(),
			ConsumerID:   payload.ConsumerID,
			RestaurantID: payload.RestaurantID,
			Items:        items,
			Address:      h.toAddressProto(payload.DeliverTo),
		}),
	)
}

func (domainHandlers[T]) toAddressProto(address domain.Address) *orderpb.Address {
	return &orderpb.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}
