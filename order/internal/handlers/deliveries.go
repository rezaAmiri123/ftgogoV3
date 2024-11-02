package handlers

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
)

func RegisterDeliveryHandlers(deliveryHandlers application.DomainEventhandlers, domainSubscriber ddd.EventSubscriber){
	domainSubscriber.Subscribe(domain.OrderCreated{}, deliveryHandlers.OnOrderCreated)
}