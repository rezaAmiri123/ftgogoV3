package order

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/es"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
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
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "order module")
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

	jsStream := jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	outboxStore := pg.NewOutboxStore("orders.outbox", svc.DB())
	inboxStore := pg.NewInboxStore("orders.inbox", svc.DB())
	inboxHandlerMiddleware := tm.NewInboxHandlerMiddleware(inboxStore)
	stream := am.RawMessageStreamWithMiddleware(
		jsStream,
		tm.NewOutboxStreamMiddleware(outboxStore),
	)
	eventStream := am.NewEventStream(reg, stream)
	replyStream := am.NewReplyStream(reg, stream)
	domainDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("orders.events", svc.DB(), reg),
		pg.NewSnapshotStore("orders.snapshots", svc.DB(), reg),
	)
	orders := es.NewAggregateRepository[*domain.Order](domain.OrderAggregate, reg, aggregateStore)
	conn, err := grpc.Dial(ctx, svc.Config().Rpc.Address())
	if err != nil {
		return err
	}

	restaurants := grpc.NewRestaurantRepository(conn)
	// consumers := grpc.NewConsumerRepository(conn)
	// kitchens := grpc.NewKitchenRepository(conn)
	// accounts := grpc.NewAccountRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, restaurants, domainDispatcher)
	app = logging.LogApplicationAccess(app, svc.Logger())

	domainEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewDomainEventHandlers(eventStream),
		"DomainEvents", svc.Logger(),
	)
	commandHandlers := logging.LogCommandHandlerAccess[ddd.Command](
		handlers.NewCommandHandlers(app),
		"Commands", svc.Logger(),
	)
	cmdMsgHandlers := am.NewCommandMessageHandler(reg, replyStream, commandHandlers)
	msgHandlerMiddleware := am.RawMessageHandlerWithMiddleware(cmdMsgHandlers, inboxHandlerMiddleware)
	// setup Driver adapters
	if err := grpc.RegisterServer(app, svc.RPC()); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlers(domainDispatcher, domainEventHandlers)
	if err := handlers.RegisterCommandHandlers(jsStream, msgHandlerMiddleware); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(jsStream, pg.NewOutboxStore("orders.outbox", svc.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := svc.Logger()
			logger.Error().Err(err).Msg("order outbox processor encountered an error")
		}
	}()
	return nil
}
