package consumer

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/handlers"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setip Driven adapters
	reg := registry.New()
	if err = registration(reg); err != nil {
		return err
	}
	if err = consumerpb.Registration(reg); err != nil {
		return err
	}
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	accounts := grpc.NewAccountRepository(conn)
	accountHandlers := logging.LogEventHandlersAccess(
		application.NewAccountHandlers(accounts),
		"Consumer",
		mono.Logger(),
	)

	jsStream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS())
	eventStream := am.NewEventStream(reg, jsStream)
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	consumers := postgres.NewConsumerReopsitory("consumer.consumers", mono.DB())

	var app application.App
	app = application.New(consumers, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup application handlers
	integrationEventHandlers := logging.LogEventHandlersAccess[ddd.AggregateEvent](
		application.NewIntegrationEventHandlers(eventStream),
		"Consumer",
		mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	handlers.RegisterAccountHandlers(accountHandlers, domainDispatcher)
	handlers.RegisterIntegrationEventHandlers(integrationEventHandlers,domainDispatcher)
	return nil
}

func registration(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// consumer events
	if err = serde.Register(domain.ConsumerRegistered{}); err != nil {
		return
	}

	return
}
