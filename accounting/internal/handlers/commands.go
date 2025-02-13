package handlers

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/accounting/accountingpb"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/errorsotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type commandHandlers struct {
	app application.App
}

func NewCommandHandlers(reg registry.Registry, app application.App, replyPubisher am.ReplyPublisher, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewCommandHandler(reg, replyPubisher, commandHandlers{
		app: app,
	}, mws...)
}

func RegisterCommandHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	cmdMsgHandler := am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) (err error) {
		return handlers.HandleMessage(ctx, msg)
	})

	_, err := subscriber.Subscribe(accountingpb.CommandChannel, cmdMsgHandler, am.MessageFilter{
		accountingpb.AuthorizeAccountCommand,
	}, am.GroupName("account-commands"))
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
