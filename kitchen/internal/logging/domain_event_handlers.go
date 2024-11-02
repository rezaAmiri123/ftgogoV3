package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application"
	"github.com/rs/zerolog"
)

type DomainEventHandlers struct {
	application.DomainEventHandlers
	logger zerolog.Logger
}

var _ application.DomainEventHandlers = (*DomainEventHandlers)(nil)

func LogDomainEventHandlerAccess(handlers application.DomainEventHandlers, logger zerolog.Logger) DomainEventHandlers {
	return DomainEventHandlers{
		DomainEventHandlers: handlers,
		logger:              logger,
	}
}

func (h DomainEventHandlers) OnTicketAccepted(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("-->kitchen.OnTicketAccepted")
	defer func() { h.logger.Info().Err(err).Msg("<--kitchen.OnTicketAccepted") }()
	return h.DomainEventHandlers.OnTicketAccepted(ctx, event)
}
