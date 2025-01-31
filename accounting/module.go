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
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/system"
	"github.com/rezaAmiri123/ftgogoV3/internal/tm"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = accountingpb.Registration(reg); err != nil {
		return err
	}
	if err = consumerpb.Registration(reg); err != nil {
		return err
	}

	jsStream := jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	outboxStore := pg.NewOutboxStore("accounting.outbox", svc.DB())
	inboxStore := pg.NewInboxStore("accounting.inbox", svc.DB())
	inboxHandlerMiddleware := tm.NewInboxHandlerMiddleware(inboxStore)
	stream := am.RawMessageStreamWithMiddleware(
		jsStream,
		tm.NewOutboxStreamMiddleware(outboxStore),
	)

	replyStream := am.NewReplyStream(reg, stream)
	accounts := postgres.NewAccountReopsitory("accounting.accounts", svc.DB())

	var app application.App
	app = application.New(accounts)
	app = logging.LogApplicationAccess(app, svc.Logger())

	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationHandlers(app),
		"IntegrationEvents", svc.Logger(),
	)
	eventMsgHandlers := am.NewEventMessageHandler(reg, integrationEventHandlers)
	eventMsgHandlerMiddleware := am.RawMessageHandlerWithMiddleware(eventMsgHandlers, inboxHandlerMiddleware)

	commandHandlers := logging.LogCommandHandlerAccess[ddd.Command](
		handlers.NewCommandHandlers(app),
		"Commands", svc.Logger(),
	)
	cmdMsgHandlers := am.NewCommandMessageHandler(reg, replyStream, commandHandlers)
	msgHandlerMiddleware := am.RawMessageHandlerWithMiddleware(cmdMsgHandlers, inboxHandlerMiddleware)

	if err = grpc.RegisterServer(app, svc.RPC()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlers(stream, eventMsgHandlerMiddleware); err != nil {
		return err
	}
	if err = handlers.RegisterCommandHandlers(stream, msgHandlerMiddleware); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(jsStream, pg.NewOutboxStore("accounting.outbox", svc.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := svc.Logger()
			logger.Error().Err(err).Msg("order outbox processor encountered an error")
		}
	}()

	return nil
}
