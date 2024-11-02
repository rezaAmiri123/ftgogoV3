package handlers

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
)

func RegisterDeliveryHandlers(deliveryHandlers application.DomainEventHandlers, domainSunscriber ddd.EventSubscriber){
	domainSunscriber.Subscribe(domain.TicketAccepted{}, deliveryHandlers.OnTicketAccepted)
}
