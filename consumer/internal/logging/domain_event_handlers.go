package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rs/zerolog"
)

type DomainEventHandlers struct {
	application.DomainEventHandlers
	logger zerolog.Logger
}

var _ application.DomainEventHandlers = (*DomainEventHandlers)(nil)

func LogDomainEventHandlersAccess(handlers application.DomainEventHandlers, logger zerolog.Logger) DomainEventHandlers {
	return DomainEventHandlers{
		DomainEventHandlers: handlers,
		logger: logger,
	}
}

func (a DomainEventHandlers) OnConsumerRegistered(ctx context.Context, event ddd.Event) (err error) {
	a.logger.Info().Msg("-->consumer.OnConsumerRegistered")
	defer func() { a.logger.Info().Err(err).Msg("<--consumer.OnConsumerRegistered") }()
	return a.DomainEventHandlers.OnConsumerRegistered(ctx, event)
}
