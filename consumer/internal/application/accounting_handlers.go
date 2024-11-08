package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type AccountHandlers[T ddd.AggregateEvent] struct {
	accounts domain.AccountRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*AccountHandlers[ddd.AggregateEvent])(nil)

func NewAccountHandlers(accounts domain.AccountRepository) *AccountHandlers[ddd.AggregateEvent] {
	return &AccountHandlers[ddd.AggregateEvent]{
		accounts: accounts,
	}
}
func (h AccountHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.ConsumerRegisteredEvent:
		return h.OnConsumerRegistered(ctx, event)
	}
	return nil
}

func (h AccountHandlers[T]) OnConsumerRegistered(ctx context.Context, event ddd.AggregateEvent) error {
	consumerRegistered := event.Payload().(*domain.ConsumerRegistered)
	return h.accounts.CreateAccount(ctx, domain.CreateAccount{
		ID:   consumerRegistered.Consumer.ID(),
		Name: consumerRegistered.Consumer.Name,
	})
}
