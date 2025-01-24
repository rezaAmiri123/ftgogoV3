package handlers

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
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

	return subscriber.Subscribe(orderpb.CommandChannel, cmdMsgHandler, am.MessageFilter{
		orderpb.ApproveOrderCommand,
	}, am.GroupName("order-commands"))
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case orderpb.ApproveOrderCommand:
		return h.doApproveOrder(ctx, cmd)
	}
	return nil, nil
}

func (h commandHandlers) doApproveOrder(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*orderpb.ApproveOrder)

	return nil, h.app.ApproveOrder(ctx, commands.ApproveOrder{
		ID:       payload.GetOrderID(),
		TicketID: payload.GetTicketID(),
	})
}
