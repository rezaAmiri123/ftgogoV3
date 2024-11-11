package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type IntegrationEventHandlers[T ddd.AggregateEvent] struct {
	publisher am.MessagePulisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*IntegrationEventHandlers[ddd.AggregateEvent])(nil)

func NewIntegrationEventHandlers(publisher am.MessagePulisher[ddd.Event]) *IntegrationEventHandlers[ddd.AggregateEvent] {
	return &IntegrationEventHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func (h IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	switch event.EventName() {
	case domain.ConsumerRegisteredEvent:
		return h.onConsumerRegistered(ctx, event)
	}
	return nil
}

func (h IntegrationEventHandlers[T]) onConsumerRegistered(ctx context.Context, event ddd.AggregateEvent) (err error) {
	payload := event.Payload().(*domain.ConsumerRegistered)
	evt := ddd.NewEvent(consumerpb.ConsumerRegisteredEvent, &consumerpb.ConsumerRegistred{
		Id:   payload.Consumer.ID(),
		Name: payload.Consumer.Name,
	})
	return h.publisher.Publish(ctx, consumerpb.ConsumerAggregateChannel, evt)
}
