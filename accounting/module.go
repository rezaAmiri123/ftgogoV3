package accounting

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/handlers"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = consumerpb.Registration(reg); err != nil {
		return err
	}
	jStream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS())
	eventStream := am.NewEventStream(reg, jStream)
	accounts := postgres.NewAccountReopsitory("accounting.accounts", mono.DB())

	var app application.App
	app = application.New(accounts)
	app = logging.LogApplicationAccess(app, mono.Logger())

	consumerHandlers := logging.LogEventHandlerAccess[ddd.Event](
		application.NewConsumerHandlers(mono.Logger()),
		"account", mono.Logger(),
	)

	if err = grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err = handlers.RegisterConsumerHandlers(consumerHandlers, eventStream); err != nil {
		return err
	}

	return nil
}
