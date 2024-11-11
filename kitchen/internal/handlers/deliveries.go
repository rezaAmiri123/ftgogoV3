package handlers

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
)

func RegisterDeliveryHandlers(deliveryHandlers ddd.EventHandler[ddd.AggregateEvent], domainSunscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSunscriber.Subscribe(deliveryHandlers, domain.TicketAcceptedEvent)
}
