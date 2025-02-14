package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rs/zerolog"
)

type EventHandlers struct {
	am.MessageHandler
	serviceName string
	label       string
	logger      zerolog.Logger
}

var _ am.MessageHandler = (*EventHandlers)(nil)

func LogMessageAccessAccess(handlers am.MessageHandler, serviceName string, label string, logger zerolog.Logger) am.MessageHandler {
	return EventHandlers{
		MessageHandler: handlers,
		serviceName:    serviceName,
		label:          label,
		logger:         logger,
	}
}

func (h EventHandlers) HandleEvent(ctx context.Context, msg am.IncomingMessage) (err error) {
	h.logger.Info().Msgf("--> %s.%s.On(%s)", h.serviceName, h.label, msg.MessageName())
	defer func() { h.logger.Info().Err(err).Msgf("<-- %s.%s.On(%s)", h.serviceName, h.label, msg.MessageName()) }()
	return h.MessageHandler.HandleMessage(ctx, msg)
}
