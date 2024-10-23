package kitchen

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
)

type Module struct{}

func(m Module)Startup(ctx context.Context, mono monolith.Monolith)error{
	tickets := postgres.NewTicketReopsitory("kitchen.tickets", mono.DB())

	var app application.App
	app = application.New(tickets)
	app = logging.LogApplicationAccess(app,mono.Logger())

	if err := grpc.RegisterServer(app,mono.RPC());err!= nil{
		return err
	}

	return nil
}
