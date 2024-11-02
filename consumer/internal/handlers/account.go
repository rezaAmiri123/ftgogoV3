package handlers

import (
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

func RegisterAccountHandlers(accountHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber){
	domainSubscriber.Subscribe(domain.ConsumerRegistered{}, accountHandlers.OnConsumerRegistered)
}