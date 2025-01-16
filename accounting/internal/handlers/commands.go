package handlers

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/accountingpb"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/application"
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

func RegisterCommandHandlers(subscriber am.CommandSubscriber, handlers ddd.CommandHandler[ddd.Command]) error {
	cmdMsgHandler := am.CommandMessageHandlerFunc(func(ctx context.Context, cmdMsg am.IncomingCommandMessage) (ddd.Reply, error) {
		return handlers.HandleCommand(ctx, cmdMsg)
	})

	return subscriber.Subscribe(accountingpb.CommandChannel, cmdMsgHandler, am.MessageFilter{
		accountingpb.AuthorizeAccountCommand,
	}, am.GroupName("account-commands"))
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case accountingpb.AuthorizeAccountCommand:
		return h.onAuthorizeAccount(ctx, cmd)
	}

	return nil, nil
}

func (h commandHandlers) onAuthorizeAccount(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*accountingpb.AuthorizeAccount)

	return nil, h.app.AuthorizeOrderByAccount(ctx, application.AuthorizeOrderByAccount{
		ID:         payload.GetAccountId(),
		OrderID:    payload.GetOrderId(),
		OrderTotal: int(payload.GetTotalOrder()),
	})
}
