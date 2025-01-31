package customerweb

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/rest"
	"github.com/rezaAmiri123/ftgogoV3/internal/system"
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

	consumers := grpc.NewConsumerRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	var app application.App
	app = application.New(consumers, orders)
	app = logging.LogApplicationAccess(app, svc.Logger())

	server := rest.NewServer(app, svc.Config().Secret)
	svc.Mux().Mount("/api/v1", server.Mount())
	svc.Mux().Mount("/spec-customer", rest.SwaggerHandler())

	return nil
}
