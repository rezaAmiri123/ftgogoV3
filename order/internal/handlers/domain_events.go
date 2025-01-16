package handlers

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
)

type domainHandlers[T ddd.Event] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.Event] = (*domainHandlers[ddd.Event])(nil)

func NewDomainEventHandlers(publisher am.MessagePublisher[ddd.Event]) ddd.EventHandler[ddd.Event] {
	return &domainHandlers[ddd.Event]{
		publisher: publisher,
	}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handlers ddd.EventHandler[ddd.Event]) {
	subscriber.Subscribe(handlers,
		domain.OrderCreatedEvent,
	)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) error {
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
