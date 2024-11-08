package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
)

type DliveryHandlers[T ddd.AggregateEvent] struct {
	deliveries domain.DeliveryRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*DliveryHandlers[ddd.AggregateEvent])(nil)

func NewDliveryHandlers(deliveries domain.DeliveryRepository)*DliveryHandlers[ddd.AggregateEvent]{
	return &DliveryHandlers[ddd.AggregateEvent]{
		deliveries: deliveries,
	}
}

func(h DliveryHandlers[T])HandleEvent(ctx context.Context, event T) error{
	switch event.EventName(){
	case domain.OrderCreatedEvent:
		return h.onOrderCreated(ctx,event)
	}
	return nil
}

func(h DliveryHandlers[T])onOrderCreated(ctx context.Context, event ddd.AggregateEvent) error{
	orderCreated := event.Payload().(*domain.OrderCreated)
	return h.deliveries.CreateDelivery(ctx,domain.CreateDelivery{
		DeliveryID: event.AggregateID(),
		RestaurantID: orderCreated.RestaurantID,
		Address: orderCreated.DeliverTo,
	})
}