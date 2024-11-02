package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type DomainEventHandlers interface {
	OnTicketAccepted(ctx context.Context, event ddd.Event) error
}
type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnTicketAccepted(ctx context.Context, event ddd.Event) error {
	return nil
}
