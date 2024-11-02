package order

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/handlers"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/postgres"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	orders := postgres.NewOrderRepository("orders.orders", mono.DB())
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
	app = application.New(orders, restaurants, consumers, kitchens, accounts, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup application handlers
	deliveryHandlers := logging.LogDomainEventHandlers(
		application.NewDliveryHandlers(deliveries),
		mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	handlers.RegisterDeliveryHandlers(deliveryHandlers, domainDispatcher)
	return nil
}
