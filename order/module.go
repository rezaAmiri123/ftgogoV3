package order

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/order/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	orders := postgres.NewOrderRepository("orders.orders", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	restaurants := grpc.NewRestaurantRepository(conn)

	var app application.App
	app = application.New(orders, restaurants)
	app = logging.LogApplicationAccess(app, mono.Logger())

	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}

	return nil
}
