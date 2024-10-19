package consumer

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
)

type Module struct{}

func(m Module)Startup(ctx context.Context, mono monolith.Monolith)error{
	consumers := postgres.NewConsumerReopsitory("consumer.consumers", mono.DB())

	var app application.App
	app = application.New(consumers)
	app = logging.LogApplicationAccess(app,mono.Logger())

	if err := grpc.RegisterServer(app,mono.RPC());err!= nil{
		return err
	}

	return nil
}
