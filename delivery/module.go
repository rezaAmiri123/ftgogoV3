package delivery

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/grpc"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/handlers"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/amotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/amprom"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/postgresotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/system"
	"github.com/rezaAmiri123/ftgogoV3/internal/tm"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	serviceName := "delivery"
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "delivery module")
		}
	}()

	// setup Driven adapters
	reg := registry.New()
	if err = orderpb.Registeration(reg); err != nil {
		return err
	}

	var stream am.MessageStream
	stream = jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	inboxStore := pg.NewInboxStore("delivery.inbox", postgresotel.Trace(svc.DB()))

	messageSubscriber := am.NewMessageSubscriber(
		stream,
		amotel.OtelMessageContextExtractor(),
		amprom.ReceivedMessagesCounter(serviceName),
	)

	deliveries := postgres.NewDeliveryRepository("delivery.deliveries", postgresotel.Trace(svc.DB()))
	couriers := postgres.NewCourierRepository("delivery.couriers", postgresotel.Trace(svc.DB()))
	conn, err := grpc.Dial(ctx, svc.Config().Rpc.Address())
	if err != nil {
		return err
	}

	restaurants := grpc.NewRestaurantRepository(conn)

	var app application.App
	app = application.New(deliveries, couriers, restaurants)
	app = logging.LogApplicationAccess(app, svc.Logger())

	if err := grpc.RegisterServer(app, svc.RPC()); err != nil {
		return err
	}

	integrationEventHandlers := am.LogMessageHandlerAccess(
		handlers.NewIntegrationHandlers(reg, app, tm.InboxHandler(inboxStore)),
		serviceName, "IntegrationEvents", svc.Logger(),
	)

	if err = handlers.RegisterIntegrationEventHandlers(messageSubscriber, integrationEventHandlers); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(stream, pg.NewOutboxStore("delivery.outbox", svc.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := svc.Logger()
			logger.Error().Err(err).Msg("order outbox processor encountered an error")
		}
	}()

	return nil
}
