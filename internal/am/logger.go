package am

import (
	"context"

	"github.com/rs/zerolog"
)

type msgHandlers struct {
	MessageHandler
	serviceName string
	label       string
	logger      zerolog.Logger
}

var _ MessageHandler = (*msgHandlers)(nil)

func LogMessageHandlerAccess(handlers MessageHandler, serviceName string, label string, logger zerolog.Logger) MessageHandler {
	return msgHandlers{
		MessageHandler: handlers,
		serviceName:    serviceName,
		label:          label,
		logger:         logger,
	}
}

func (h msgHandlers) HandleEvent(ctx context.Context, msg IncomingMessage) (err error) {
	h.logger.Info().Msgf("--> %s.%s.On(%s)", h.serviceName, h.label, msg.MessageName())
	defer func() { h.logger.Info().Err(err).Msgf("<-- %s.%s.On(%s)", h.serviceName, h.label, msg.MessageName()) }()
	return h.MessageHandler.HandleMessage(ctx, msg)
}
