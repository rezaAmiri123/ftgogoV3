package order

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/es"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	"github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/handlers"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/logging"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	reg := registry.New()
	err := registerations(reg)
	if err != nil {
		return err
	}

	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	// orders := postgres.NewOrderRepository("orders.orders", mono.DB())
	aggregateStore := es.AggregateStoreWithMiddleware(
		postgres.NewEventStore("orders.events", mono.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		postgres.NewSnapshotStore("orders.snapshots", mono.DB(), reg),
	)
	orders := es.NewAggregateRepository[*domain.Order](domain.OrderAggregate, reg, aggregateStore)
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	restaurants := grpc.NewRestaurantRepository(conn)
	consumers := grpc.NewConsumerRepository(conn)
	kitchens := grpc.NewKitchenRepository(conn)
	accounts := grpc.NewAccountRepository(conn)
	deliveries := grpc.NewDeliveryRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, restaurants, consumers, kitchens, accounts)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup application handlers
	deliveryHandlers := logging.LogEventHandlerAccess(
		application.NewDliveryHandlers(deliveries),
		"order",
		mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	handlers.RegisterDeliveryHandlers(deliveryHandlers, domainDispatcher)
	return nil
}

func registerations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Order
	err = serde.Register(domain.Order{}, func(v any) error {
		order := v.(*domain.Order)
		order.Aggregate = es.NewAggregate("", domain.OrderAggregate)
		order.Status = domain.UnknownOrderStatus
		return nil
	})
	if err != nil {
		return err
	}

	// order events
	if err = serde.Register(domain.OrderCreated{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderApproved{}); err != nil {
		return err
	}

	return nil
}
