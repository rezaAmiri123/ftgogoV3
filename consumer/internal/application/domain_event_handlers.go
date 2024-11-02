package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type DomainEventHandlers interface {
	OnConsumerRegistered(ctx context.Context, event ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnConsumerRegistered(ctx context.Context, event ddd.Event) error {
	return nil
}
