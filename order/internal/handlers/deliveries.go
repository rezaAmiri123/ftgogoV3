package handlers

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
)

func RegisterDeliveryHandlers(deliveryHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]){
	domainSubscriber.Subscribe(domain.OrderCreatedEvent, deliveryHandlers)
}