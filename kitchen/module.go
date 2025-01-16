package kitchen

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/handlers"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"github.com/stackus/errors"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "order module")
		}
	}()

	// setup Driven adapters
	reg := registry.New()
	if err = kitchenpb.Registeration(reg); err != nil {
		return err
	}

	stream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())
	commandStream := am.NewCommandStream(reg, stream)

	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	tickets := postgres.NewTicketReopsitory("kitchen.tickets", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	deliveries := grpc.NewDeliveryRepository(conn)

	// setup application
	var app application.App
	app = application.New(tickets, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup application handlers
	deliveryHandlers := logging.LogEventHandlerAccess(
		application.NewDeliveryHandlers(deliveries),
		"delivery",
		mono.Logger(),
	)

	commandHandlers := logging.LogCommandHandlersAccess[ddd.Command](
		handlers.NewCommandHandlers(app),
		"Commands", mono.Logger(),
	)
	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	handlers.RegisterDeliveryHandlers(deliveryHandlers, domainDispatcher)

	if err = handlers.RegisterCommandHandlers(commandStream, commandHandlers); err != nil {
		return err
	}
	
	return nil
}
