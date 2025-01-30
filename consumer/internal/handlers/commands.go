package handlers

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type commandHandlers struct {
	app application.App
}

func NewCommandHandlers(app application.App) ddd.CommandHandler[ddd.Command] {
	return commandHandlers{
		app: app,
	}
}

func RegisterCommandHandlers(subscriber am.RawMessageStream, handlers am.RawMessageHandler) error {
	cmdMsgHandler := am.RawMessageHandlerFunc(func(ctx context.Context, msg am.IncomingRawMessage) (err error) {
		return handlers.HandleMessage(ctx, msg)
	})

	_, err := subscriber.Subscribe(consumerpb.CommandChannel, cmdMsgHandler, am.MessageFilter{
		consumerpb.AuthorizeConsumerCommand,
	}, am.GroupName("consumer-commands"))
	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case consumerpb.AuthorizeConsumerCommand:
		return h.onAuthorizeConsumer(ctx, cmd)
	}

	return nil, nil
}

func (h commandHandlers) onAuthorizeConsumer(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*consumerpb.AuthorizeCustomer)

	return nil, h.app.ValidateOrderByConsumer(ctx, application.ValidateOrderByConsumer{
		ConsumerID: payload.GetId(),
		OrderID:    payload.GetOrderId(),
		OrderTotal: int(payload.GetTotalOrder()),
	})
}
