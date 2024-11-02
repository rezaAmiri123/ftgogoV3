package delivery

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	deliveries := postgres.NewDeliveryRepository("delivery.deliveries", mono.DB())
	couriers := postgres.NewCourierRepository("delivery.couriers", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	restaurants := grpc.NewRestaurantRepository(conn)

	var app application.App
	app = application.New(deliveries, couriers, restaurants)
	app = logging.LogApplicationAccess(app, mono.Logger())

	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}

	return nil
}
