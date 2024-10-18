package accounting

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
)

type Module struct{}

func(m Module)Startup(ctx context.Context, mono monolith.Monolith)error{
	accounts := postgres.NewAccountReopsitory("accounting.accounts", mono.DB())

	var app application.App
	app = application.New(accounts)
	app = logging.LogApplicationAccess(app,mono.Logger())

	if err := grpc.RegisterServer(app,mono.RPC());err!= nil{
		return err
	}

	return nil
}
