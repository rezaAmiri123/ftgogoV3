package handlers

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

func RegisterConsumerHandlers(consumerHandlers ddd.EventHandler[ddd.Event], subscriber am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return consumerHandlers.HandleEvent(ctx, eventMsg)
	})

	msgFilters := am.MessageFilter{
		consumerpb.ConsumerRegisteredEvent,
	}

	return subscriber.Subscribe(consumerpb.ConsumerAggregateChannel, evtMsgHandler, msgFilters)

}
