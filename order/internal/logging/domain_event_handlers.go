package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application"
	"github.com/rs/zerolog"
)

type DomainEventHandlers struct{
	application.DomainEventhandlers
	logger zerolog.Logger
}

var _ application.DomainEventhandlers = (*DomainEventHandlers)(nil)

func LogDomainEventHandlers(handlers application.DomainEventhandlers, logger zerolog.Logger)DomainEventHandlers{
	return DomainEventHandlers{
		DomainEventhandlers: handlers,
		logger: logger,
	}
}

func(h DomainEventHandlers)OnOrderCreated(ctx context.Context, event ddd.Event) (err error){
	h.logger.Info().Msg("-->order.OnOrderCreated")
	defer func() { h.logger.Info().Err(err).Msg("<--order.OnOrderCreated") }()
	return h.DomainEventhandlers.OnOrderCreated(ctx, event)
}