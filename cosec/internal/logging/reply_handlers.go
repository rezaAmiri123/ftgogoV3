package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/sec"
	"github.com/rs/zerolog"
)

type sagaReplyHandlers[T any] struct {
	sec.Orchestrator[T]
	label  string
	logger zerolog.Logger
}

var _ sec.Orchestrator[any] = (*sagaReplyHandlers[any])(nil)

func LogReplyHandlersAccess[T any](orc sec.Orchestrator[T], label string, logger zerolog.Logger) sec.Orchestrator[T] {
	return sagaReplyHandlers[T]{
		Orchestrator: orc,
		label:        label,
		logger:       logger,
	}
}

func (h sagaReplyHandlers[T]) HandleReply(ctx context.Context, reply ddd.Reply) error {
	h.logger.Info().Msgf("---> COSEC.%s.On(%s)", h.label, reply.ReplyName())
	defer func() { h.logger.Info().Msgf("<--- COSEC.%s.On(%s)", h.label, reply.ReplyName()) }()
	return h.Orchestrator.HandleReply(ctx, reply)
}
