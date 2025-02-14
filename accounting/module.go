package accounting

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/accountingpb"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/handlers"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/amotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/amprom"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/postgresotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/system"
	"github.com/rezaAmiri123/ftgogoV3/internal/tm"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	serviceName := "accounting"
	// setup Driven adapters
	reg := registry.New()
	if err = accountingpb.Registration(reg); err != nil {
		return err
	}
	if err = consumerpb.Registration(reg); err != nil {
		return err
	}

	var stream am.MessageStream
	stream = jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	outboxStore := pg.NewOutboxStore("accounting.outbox", postgresotel.Trace(svc.DB()))
	inboxStore := pg.NewInboxStore("accounting.inbox", postgresotel.Trace(svc.DB()))

	sentCounter := amprom.SentMessageCounter(serviceName)
	messagePublisher := am.NewMessagePublisher(
		stream,
		amotel.OtelMessageContextInjector(),
		sentCounter,
		tm.OutboxPublisher(outboxStore),
	)
	messageSubscriber := am.NewMessageSubscriber(
		stream,
		amotel.OtelMessageContextExtractor(),
		amprom.ReceivedMessagesCounter(serviceName),
	)
	replyPublisher := am.NewReplyPublisher(reg, messagePublisher)

	accounts := postgres.NewAccountReopsitory("accounting.accounts", postgresotel.Trace(svc.DB()))

	var app application.App
	app = application.New(accounts)
	app = logging.LogApplicationAccess(app, svc.Logger())

	integrationEventHandlers := am.LogMessageHandlerAccess(
		handlers.NewIntegrationHandlers(reg, app, tm.InboxHandler(inboxStore)),
		serviceName, "IntegrationEvents", svc.Logger(),
	)

	commandHandlers := am.LogMessageHandlerAccess(
		handlers.NewCommandHandlers(reg, app, replyPublisher, tm.InboxHandler(inboxStore)),
		serviceName, "Commands", svc.Logger(),
	)

	if err = grpc.RegisterServer(app, svc.RPC()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlers(messageSubscriber, integrationEventHandlers); err != nil {
		return err
	}
	if err = handlers.RegisterCommandHandlers(messageSubscriber, commandHandlers); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(stream, pg.NewOutboxStore("accounting.outbox", svc.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := svc.Logger()
			logger.Error().Err(err).Msg("order outbox processor encountered an error")
		}
	}()

	return nil
}
