package order

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/amotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/amprom"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/es"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/postgresotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/system"
	"github.com/rezaAmiri123/ftgogoV3/internal/tm"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/handlers"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	serviceName := "orders"
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "orders module")
		}
	}()

	// setup Driven adapters
	reg := registry.New()
	if err = domain.Registerations(reg); err != nil {
		return err
	}
	if err = orderpb.Registeration(reg); err != nil {
		return err
	}

	var stream am.MessageStream
	stream = jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	outboxStore := pg.NewOutboxStore("orders.outbox", postgresotel.Trace(svc.DB()))
	inboxStore := pg.NewInboxStore("orders.inbox", postgresotel.Trace(svc.DB()))

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

	domainDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("orders.events", postgresotel.Trace(svc.DB()), reg),
		pg.NewSnapshotStore("orders.snapshots", postgresotel.Trace(svc.DB()), reg),
	)
	orders := es.NewAggregateRepository[*domain.Order](
		domain.OrderAggregate, 
		reg, 
		aggregateStore,
	)
	conn, err := grpc.Dial(ctx, svc.Config().Rpc.Address())
	if err != nil {
		return err
	}

	restaurants := grpc.NewRestaurantRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, restaurants, domainDispatcher)
	app = logging.LogApplicationAccess(app, svc.Logger())

	domainEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
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
	handlers.RegisterDomainEventHandlers(domainDispatcher, domainEventHandlers)
	if err = handlers.RegisterCommandHandlers(messageSubscriber, commandHandlers); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(stream, pg.NewOutboxStore("orders.outbox", svc.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := svc.Logger()
			logger.Error().Err(err).Msg("order outbox processor encountered an error")
		}
	}()
	return nil
}
