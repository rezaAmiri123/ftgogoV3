package cosec

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/accountingpb"
	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/cosec/internal"
	"github.com/rezaAmiri123/ftgogoV3/cosec/internal/handlers"
	"github.com/rezaAmiri123/ftgogoV3/cosec/internal/logging"
	"github.com/rezaAmiri123/ftgogoV3/cosec/internal/models"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/amotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/amprom"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	"github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/postgresotel"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
	"github.com/rezaAmiri123/ftgogoV3/internal/sec"
	"github.com/rezaAmiri123/ftgogoV3/internal/system"
	"github.com/rezaAmiri123/ftgogoV3/internal/tm"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
	"github.com/stackus/errors"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	serviceName := "cosec"
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "cosec module")
		}
	}()

	// setup Driven adapters
	reg := registry.New()
	if err = registerations(reg); err != nil {
		return err
	}
	if err = orderpb.Registeration(reg); err != nil {
		return err
	}
	if err = consumerpb.Registration(reg); err != nil {
		return err
	}
	if err = kitchenpb.Registeration(reg); err != nil {
		return err
	}
	if err = accountingpb.Registration(reg); err != nil {
		return err
	}

	var stream am.MessageStream
	stream = jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	outboxStore := pg.NewOutboxStore("cosec.outbox", postgresotel.Trace(svc.DB()))
	inboxStore := pg.NewInboxStore("cosec.inbox", postgresotel.Trace(svc.DB()))

	sentCounter := amprom.SentMessageCounter(serviceName)
	messagePublisher := am.NewMessagePublisher(
		stream,
		amotel.OtelMessageContextInjector(),
		sentCounter,
		tm.OutboxPublisher(outboxStore),
	)
	messageSubscriber := am.NewMessageSubscriber(
		stream,
		amotel.OtelMessageContextExtractor(),
		amprom.ReceivedMessagesCounter(serviceName),
	)

	commandPublisher := am.NewCommandPublisher(reg, messagePublisher)
	sagaStore := postgres.NewSagaStore("cosec.sagas", postgresotel.Trace(svc.DB()), reg)
	sagaRepo := sec.NewSagaRepository[*models.CreateOrderData](reg, sagaStore)

	// setup application
	orchestrator := logging.LogReplyHandlersAccess[*models.CreateOrderData](
		sec.NewOrchestrator[*models.CreateOrderData](internal.NewCreateOrderSaga(), sagaRepo, commandPublisher),
		"CreateOrderSaga", svc.Logger(),
	)
	integrationEventHandlers := am.LogMessageHandlerAccess(
		handlers.NewIntegrationEventHandlers(reg, orchestrator, tm.InboxHandler(inboxStore)),
		serviceName, "IntegrationEvents", svc.Logger(),
	)

	// setup driver adapters
	if err = handlers.RegisterIntegrationEventHandlers(messageSubscriber, integrationEventHandlers); err != nil {
		return err
	}

	// replyHandler := handlers.
	replyHandler := am.NewReplyHandler(reg, orchestrator, tm.InboxHandler(inboxStore))
	if err = handlers.RegisterReplyHandlers(messageSubscriber, replyHandler); err != nil {
		return err
	}

	outboxProcessor := tm.NewOutboxProcessor(stream, pg.NewOutboxStore("cosec.outbox", svc.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := svc.Logger()
			logger.Error().Err(err).Msg("order outbox processor encountered an error")
		}
	}()

	return nil
}

func registerations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Saga data
	if err = serde.RegisterKey(internal.CreateOrderSagaName, models.CreateOrderData{}); err != nil {
		return err
	}
	return nil
}
