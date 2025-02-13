package consumer

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/handlers"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/amotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/amprom"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/postgresotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
	"github.com/rezaAmiri123/ftgogoV3/internal/system"
	"github.com/rezaAmiri123/ftgogoV3/internal/tm"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	serviceName := "consumer"
	// setip Driven adapters
	reg := registry.New()
	if err = registration(reg); err != nil {
		return err
	}
	if err = consumerpb.Registration(reg); err != nil {
		return err
	}
	var stream am.MessageStream
	stream = jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	outboxStore := pg.NewOutboxStore("consumer.outbox", postgresotel.Trace(svc.DB()))
	inboxStore := pg.NewInboxStore("consumer.inbox", postgresotel.Trace(svc.DB()))

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
	eventPublisher := am.NewEventPublisher(reg, messagePublisher)
	replyPublisher := am.NewReplyPublisher(reg, messagePublisher)

	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	consumers := postgres.NewConsumerReopsitory("consumer.consumers", postgresotel.Trace(svc.DB()))

	var app application.App
	app = application.New(consumers, domainDispatcher)
	app = logging.LogApplicationAccess(app, svc.Logger())

	// setup application handlers
	domainEventHandlers := logging.LogEventHandlersAccess[ddd.AggregateEvent](
		handlers.NewDomainEventHandlers(eventPublisher),
		"DomainEvents", svc.Logger(),
	)
	commandHandlers := am.LogMessageHandlerAccess(
		handlers.NewCommandHandlers(reg, app, replyPublisher, tm.InboxHandler(inboxStore)),
		serviceName, "Commands", svc.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(app, svc.RPC()); err != nil {
		return err
	}
	// handlers.RegisterAccountHandlers(accountHandlers, domainDispatcher)
	handlers.RegisterDomainEventHandlers(domainDispatcher, domainEventHandlers)
	if err = handlers.RegisterCommandHandlers(messageSubscriber, commandHandlers); err != nil {
		return err
	}
	if err = consumerpb.RegisterAsyncAPI(svc.Mux()); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(stream, pg.NewOutboxStore("consumer.outbox", svc.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := svc.Logger()
			logger.Error().Err(err).Msg("order outbox processor encountered an error")
		}
	}()

	return nil
}

func registration(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// consumer events
	if err = serde.Register(domain.ConsumerRegistered{}); err != nil {
		return
	}

	return
}
