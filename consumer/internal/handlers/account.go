package handlers

import (
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

func RegisterAccountHandlers(accountHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]){
	domainSubscriber.Subscribe(accountHandlers, domain.ConsumerRegisteredEvent )
}