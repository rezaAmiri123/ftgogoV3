package kitchen

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/tm"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/handlers"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"github.com/stackus/errors"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "order module")
		}
	}()

	// setup Driven adapters
	reg := registry.New()
	if err = kitchenpb.Registeration(reg); err != nil {
		return err
	}

	jsStream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())
	outboxStore := pg.NewOutboxStore("kitchen.outbox", mono.DB())
	inboxStore := pg.NewInboxStore("kitchen.inbox", mono.DB())
	inboxHandlerMiddleware := tm.NewInboxHandlerMiddleware(inboxStore)
	stream := am.RawMessageStreamWithMiddleware(
		jsStream,
		tm.NewOutboxStreamMiddleware(outboxStore),
	)
	replyStream := am.NewReplyStream(reg, stream)

	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	tickets := postgres.NewTicketReopsitory("kitchen.tickets", mono.DB())

	// setup application
	var app application.App
	app = application.New(tickets, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	commandHandlers := logging.LogCommandHandlersAccess[ddd.Command](
		handlers.NewCommandHandlers(app),
		"Commands", mono.Logger(),
	)
	cmdMsgHandlers := am.NewCommandMessageHandler(reg, replyStream, commandHandlers)
	msgHandlerMiddleware := am.RawMessageHandlerWithMiddleware(cmdMsgHandlers, inboxHandlerMiddleware)
	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}

	if err = handlers.RegisterCommandHandlers(stream, msgHandlerMiddleware); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(jsStream, pg.NewOutboxStore("kitchen.outbox", mono.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := mono.Logger()
			logger.Error().Err(err).Msg("order outbox processor encountered an error")
		}
	}()

	return nil
}
