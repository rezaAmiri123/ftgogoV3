package restaurant

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/system"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/postgres"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	consumers := postgres.NewRestaurantRepository("restaurant.restaurants", svc.DB())

	var app application.App
	app = application.New(consumers)
	app = logging.LogApplicationAccess(app, svc.Logger())

	if err := grpc.RegisterServer(app, svc.RPC()); err != nil {
		return err
	}

	return nil
}
