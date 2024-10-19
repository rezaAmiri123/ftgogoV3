package restaurant

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/postgres"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	consumers := postgres.NewRestaurantRepository("restaurant.restaurants", mono.DB())

	var app application.App
	app = application.New(consumers)
	app = logging.LogApplicationAccess(app, mono.Logger())

	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}

	return nil
}
