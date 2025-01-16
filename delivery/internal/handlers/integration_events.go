package handlers

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
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

func RegisterIntegrationEventHandlers(subscriber am.EventSubscriber, handlers ddd.EventHandler[ddd.Event]) (err error) {
	evtMsgHandler := am.MessageHandlerFunc[am.IncomingEventMessage](func(ctx context.Context, msg am.IncomingEventMessage) error {
		return handlers.HandleEvent(ctx, msg)
	})

	err = subscriber.Subscribe(kitchenpb.TicketAggregateChannel, evtMsgHandler, am.MessageFilter{
		kitchenpb.TicketAcceptedEvent,
	}, am.GroupName("delivery-tickets"))
	if err != nil {
		return err
	}
	return subscriber.Subscribe(orderpb.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderpb.OrderCreatedEvent,
	}, am.GroupName("delivery-orders"))
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) error {
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
		ID:           payload.GetOrderID(),
		RestaurantID: payload.GetRestaurantID(),
		DeliveryAddress: h.toAddressDomain(payload.GetAddress()),
	})
}

func (integrationHandlers[T]) toAddressDomain(address *orderpb.Address) domain.Address {
    if address == nil{
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
