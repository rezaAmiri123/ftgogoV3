package kitchen

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/amotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/amprom"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/postgresotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/system"
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

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	serviceName := "kitchen"
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "kitchen module")
		}
	}()

	// setup Driven adapters
	reg := registry.New()
	if err = kitchenpb.Registeration(reg); err != nil {
		return err
	}

	var stream am.MessageStream
	stream = jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	outboxStore := pg.NewOutboxStore("kitchen.outbox", postgresotel.Trace(svc.DB()))
	inboxStore := pg.NewInboxStore("kitchen.inbox", postgresotel.Trace(svc.DB()))

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

	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	tickets := postgres.NewTicketReopsitory("kitchen.tickets", postgresotel.Trace(svc.DB()))

	// setup application
	var app application.App
	app = application.New(tickets, domainDispatcher)
	app = logging.LogApplicationAccess(app, svc.Logger())

	commandHandlers := am.LogMessageHandlerAccess(
		handlers.NewCommandHandlers(reg, app, replyPublisher, tm.InboxHandler(inboxStore)),
		serviceName, "Commands", svc.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(app, svc.RPC()); err != nil {
		return err
	}

	if err = handlers.RegisterCommandHandlers(messageSubscriber, commandHandlers); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(stream, pg.NewOutboxStore("kitchen.outbox", svc.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := svc.Logger()
			logger.Error().Err(err).Msg("order outbox processor encountered an error")
		}
	}()

	return nil
}
