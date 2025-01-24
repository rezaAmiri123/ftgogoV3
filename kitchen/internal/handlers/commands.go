package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
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

	return subscriber.Subscribe(kitchenpb.CommandChannel, cmdMsgHandler, am.MessageFilter{
		kitchenpb.CreateTicketCommands,
		kitchenpb.ConfirmCreateTicketCommands,
	}, am.GroupName("kitchen-commands"))
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case kitchenpb.CreateTicketCommands:
		return h.doCreateTicket(ctx, cmd)
	case kitchenpb.ConfirmCreateTicketCommands:
		return h.doConfirmCreateTicket(ctx, cmd)

	}
	return nil, nil
}

func (h commandHandlers) doCreateTicket(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*kitchenpb.CreateTicket)

	id := uuid.New().String()

	items := make([]domain.LineItem, len(payload.Items))
	for i, item := range payload.Items {
		items[i] = domain.LineItem{
			MenuItemID: item.GetMenuItemID(),
			Name:       item.GetName(),
			Quantity:   int(item.GetQuantity()),
		}
	}
	err := h.app.CreateTicket(ctx, commands.CreateTicket{
		ID:           id,
		RestaurantID: payload.GetRestaurantID(),
		LineItems:    items,
	})

	return ddd.NewReply(kitchenpb.CreatedTicketReply, &kitchenpb.CreatedTicket{Id: id}), err
}

func (h commandHandlers) doConfirmCreateTicket(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*kitchenpb.ConfirmCreateTicket)

	return nil, h.app.ConfirmCreateTicket(ctx, commands.ConfirmCreateTicket{
		ID: payload.GetTicketID(),
	})
}
