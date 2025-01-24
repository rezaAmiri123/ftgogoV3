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
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
	"github.com/rezaAmiri123/ftgogoV3/internal/tm"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setip Driven adapters
	reg := registry.New()
	if err = registration(reg); err != nil {
		return err
	}
	if err = consumerpb.Registration(reg); err != nil {
		return err
	}

	jsStream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())
	outboxStore := pg.NewOutboxStore("consumer.outbox", mono.DB())
	inboxStore := pg.NewInboxStore("consumer.inbox", mono.DB())
	inboxHandlerMiddleware := tm.NewInboxHandlerMiddleware(inboxStore)
	stream := am.RawMessageStreamWithMiddleware(
		jsStream,
		tm.NewOutboxStreamMiddleware(outboxStore),
	)
	eventStream := am.NewEventStream(reg, stream)
	replyStream := am.NewReplyStream(reg, stream)
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	consumers := postgres.NewConsumerReopsitory("consumer.consumers", mono.DB())

	var app application.App
	app = application.New(consumers, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup application handlers
	domainEventHandlers := logging.LogEventHandlersAccess[ddd.AggregateEvent](
		handlers.NewDomainEventHandlers(eventStream),
		"DomainEvents", mono.Logger(),
	)
	commandHandlers := logging.LogCommandHandlerAccess[ddd.Command](
		handlers.NewCommandHandlers(app),
		"Commands", mono.Logger(),
	)
	cmdMsgHandlers := am.NewCommandMessageHandler(reg, replyStream, commandHandlers)
	msgHandlerMiddleware := am.RawMessageHandlerWithMiddleware(cmdMsgHandlers, inboxHandlerMiddleware)

	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	// handlers.RegisterAccountHandlers(accountHandlers, domainDispatcher)
	handlers.RegisterDomainEventHandlers(domainDispatcher, domainEventHandlers)
	if err = handlers.RegisterCommandHandlers(stream, msgHandlerMiddleware); err != nil {
		return err
	}
	if err = consumerpb.RegisterAsyncAPI(mono.Mux()); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(jsStream, pg.NewOutboxStore("consumer.outbox", mono.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := mono.Logger()
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
