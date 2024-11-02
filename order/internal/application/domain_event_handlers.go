package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type DomainEventhandlers interface {
	OnOrderCreated(ctx context.Context, event ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventhandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnOrderCreated(ctx context.Context, event ddd.Event) error {
	return nil
}
