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
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/tm"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = accountingpb.Registration(reg); err != nil {
		return err
	}
	if err = consumerpb.Registration(reg); err != nil {
		return err
	}

	jsStream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())
	outboxStore := pg.NewOutboxStore("accounting.outbox", mono.DB())
	inboxStore := pg.NewInboxStore("accounting.inbox", mono.DB())
	inboxHandlerMiddleware := tm.NewInboxHandlerMiddleware(inboxStore)
	stream := am.RawMessageStreamWithMiddleware(
		jsStream,
		tm.NewOutboxStreamMiddleware(outboxStore),
	)

	replyStream := am.NewReplyStream(reg, stream)
	accounts := postgres.NewAccountReopsitory("accounting.accounts", mono.DB())

	var app application.App
	app = application.New(accounts)
	app = logging.LogApplicationAccess(app, mono.Logger())

	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationHandlers(app),
		"IntegrationEvents", mono.Logger(),
	)
	eventMsgHandlers := am.NewEventMessageHandler(reg, integrationEventHandlers)
	eventMsgHandlerMiddleware := am.RawMessageHandlerWithMiddleware(eventMsgHandlers, inboxHandlerMiddleware)

	commandHandlers := logging.LogCommandHandlerAccess[ddd.Command](
		handlers.NewCommandHandlers(app),
		"Commands", mono.Logger(),
	)
	cmdMsgHandlers := am.NewCommandMessageHandler(reg, replyStream, commandHandlers)
	msgHandlerMiddleware := am.RawMessageHandlerWithMiddleware(cmdMsgHandlers, inboxHandlerMiddleware)

	if err = grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlers(stream, eventMsgHandlerMiddleware); err != nil {
		return err
	}
	if err = handlers.RegisterCommandHandlers(stream, msgHandlerMiddleware); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(jsStream, pg.NewOutboxStore("accounting.outbox", mono.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := mono.Logger()
			logger.Error().Err(err).Msg("order outbox processor encountered an error")
		}
	}()

	return nil
}
