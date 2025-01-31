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
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
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

	jsStream := jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	outboxStore := pg.NewOutboxStore("delivery.outbox", svc.DB())
	inboxStore := pg.NewInboxStore("delivery.inbox", svc.DB())
	inboxHandlerMiddleware := tm.NewInboxHandlerMiddleware(inboxStore)
	stream := am.RawMessageStreamWithMiddleware(
		jsStream,
		tm.NewOutboxStreamMiddleware(outboxStore),
	)

	deliveries := postgres.NewDeliveryRepository("delivery.deliveries", svc.DB())
	couriers := postgres.NewCourierRepository("delivery.couriers", svc.DB())
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

	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationHandlers(app),
		"IntegrationEvents", svc.Logger(),
	)
	eventMsgHandlers := am.NewEventMessageHandler(reg, integrationEventHandlers)
	msgHandlerMiddleware := am.RawMessageHandlerWithMiddleware(eventMsgHandlers, inboxHandlerMiddleware)

	if err = handlers.RegisterIntegrationEventHandlers(stream, msgHandlerMiddleware); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(jsStream, pg.NewOutboxStore("delivery.outbox", svc.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := svc.Logger()
			logger.Error().Err(err).Msg("order outbox processor encountered an error")
		}
	}()

	return nil
}
