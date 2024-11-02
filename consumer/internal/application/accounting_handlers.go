package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type AccountHandlers struct {
	accounts domain.AccountRepository
	ignoreUnimplementedDomainEvents
}

var _ DomainEventHandlers = (*AccountHandlers)(nil)

func NewAccountHandlers(accounts domain.AccountRepository) *AccountHandlers {
	return &AccountHandlers{
		accounts: accounts,
	}
}

func (h AccountHandlers) OnConsumerRegistered(ctx context.Context, event ddd.Event) error {
	consumerRegistered := event.(*domain.ConsumerRegistered)
	return h.accounts.CreateAccount(ctx, domain.CreateAccount{
		ID:   consumerRegistered.Consumer.ID,
		Name: consumerRegistered.Consumer.Name,
	})
}
