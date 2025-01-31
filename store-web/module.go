package storeweb

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/system"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/rest"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	conn, err := grpc.Dial(ctx, svc.Config().Rpc.Address())
	if err != nil {
		return err
	}

	restaurants := grpc.NewRestaurantRepository(conn)
	kitchens := grpc.NewKitchenRepository(conn)

	var app application.App
	app = application.New(restaurants, kitchens)
	app = logging.LogApplicationAccess(app, svc.Logger())

	server := rest.NewServer(app, svc.Config().Secret)
	svc.Mux().Mount("/store/v1", server.Mount())
	svc.Mux().Mount("/spec-store", rest.SwaggerHandler())

	return nil
}
