package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rs/zerolog"
)

type CommandHandlers[T ddd.Command] struct {
	ddd.CommandHandler[T]
	label  string
	logger zerolog.Logger
}

var _ ddd.CommandHandler[ddd.Command] = (*CommandHandlers[ddd.Command])(nil)

func LogCommandHandlerAccess[T ddd.Command](handlers ddd.CommandHandler[T], label string, logger zerolog.Logger) ddd.CommandHandler[T] {
	return CommandHandlers[T]{
		CommandHandler: handlers,
		label:          label,
		logger:         logger,
	}
}

func (h CommandHandlers[T]) HandleCommand(ctx context.Context, cmd T) (ddd.Reply, error) {
	h.logger.Info().Msgf("-->Orders.%s.On(%s)", h.label, cmd.CommandName())
	defer func() { h.logger.Info().Msgf("<--Orders.%s.On(%s)", h.label, cmd.CommandName()) }()
	return h.CommandHandler.HandleCommand(ctx, cmd)
}
