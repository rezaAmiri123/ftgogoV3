package customerweb

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/rest"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	consumers := grpc.NewConsumerRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	var app application.App
	app = application.New(consumers, orders)
	app = logging.LogApplicationAccess(app, mono.Logger())

	server := rest.NewServer(app, mono.Config().Secret)
	mono.Mux().Mount("/api/v1", server.Mount())
	mono.Mux().Mount("/spec-customer", rest.SwaggerHandler())

	return nil
}
