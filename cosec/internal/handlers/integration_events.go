package handlers

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/cosec/internal/models"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/sec"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
)

type integrationHandlers[T ddd.Event] struct {
	orchestrator sec.Orchestrator[*models.CreateOrderData]
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(orchestrator sec.Orchestrator[*models.CreateOrderData]) ddd.EventHandler[ddd.Event] {
	return integrationHandlers[ddd.Event]{
		orchestrator: orchestrator,
	}
}

func RegisterIntegrationEventHandlers(subscriber am.RawMessageStream, handlers am.RawMessageHandler) (err error) {
	evtMsgHandler := am.RawMessageHandlerFunc(func(ctx context.Context, msg am.IncomingRawMessage) error {
		return handlers.HandleMessage(ctx, msg)
	})
	
	return subscriber.Subscribe(orderpb.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderpb.OrderCreatedEvent,
	}, am.GroupName("cosec-orders"))
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) error {
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
