package handlers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/errorsotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type commandHandlers struct {
	app application.App
}

func NewCommandHandlers(reg registry.Registry, app application.App, replyPubliher am.ReplyPublisher, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewCommandHandler(reg, replyPubliher, commandHandlers{
		app: app,
	}, mws...)
}

func RegisterCommandHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	cmdMsgHandler := am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) (err error) {
		return handlers.HandleMessage(ctx, msg)
	})

	_, err := subscriber.Subscribe(kitchenpb.CommandChannel, cmdMsgHandler, am.MessageFilter{
		kitchenpb.CreateTicketCommands,
		kitchenpb.ConfirmCreateTicketCommands,
	}, am.GroupName("kitchen-commands"))
	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (reply ddd.Reply, err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling command",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled command", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling command", trace.WithAttributes(
		attribute.String("Command", cmd.CommandName()),
	))

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
