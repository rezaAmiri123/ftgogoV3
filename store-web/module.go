package storeweb

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/rest"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	restaurants := grpc.NewRestaurantRepository(conn)
	kitchens := grpc.NewKitchenRepository(conn)

	var app application.App
	app = application.New(restaurants, kitchens)
	app = logging.LogApplicationAccess(app, mono.Logger())

	server := rest.NewServer(app, mono.Config().Secret)
	mono.Mux().Mount("/store/v1", server.Mount())
	mono.Mux().Mount("/spec-store", rest.SwaggerHandler())

	return nil
}
