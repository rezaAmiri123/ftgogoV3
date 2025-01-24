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
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/jetstream"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	"github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	pg "github.com/rezaAmiri123/ftgogoV3/internal/postgres"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
	"github.com/rezaAmiri123/ftgogoV3/internal/sec"
	"github.com/rezaAmiri123/ftgogoV3/internal/tm"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
	"github.com/stackus/errors"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
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

	jsStream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())
	outboxStore := pg.NewOutboxStore("cosec.outbox", mono.DB())
	inboxStore := pg.NewInboxStore("cosec.inbox", mono.DB())
	inboxHandlerMiddleware := tm.NewInboxHandlerMiddleware(inboxStore)
	stream := am.RawMessageStreamWithMiddleware(
		jsStream,
		tm.NewOutboxStreamMiddleware(outboxStore),
	)
	commandStream := am.NewCommandStream(reg, stream)
	sagaStore := postgres.NewSagaStore("cosec.sagas", mono.DB(), reg)
	sagaRepo := sec.NewSagaRepository[*models.CreateOrderData](reg, sagaStore)

	// setup application
	orchestrator := logging.LogReplyHandlersAccess[*models.CreateOrderData](
		sec.NewOrchestrator[*models.CreateOrderData](internal.NewCreateOrderSaga(), sagaRepo, commandStream),
		"CreateOrderSaga", mono.Logger(),
	)

	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationEventHandlers(orchestrator),
		"IntegrationEvents", mono.Logger(),
	)
	evtMsgHandlers := am.NewEventMessageHandler(reg, integrationEventHandlers)
	msgEvtHandlerMiddleware := am.RawMessageHandlerWithMiddleware(evtMsgHandlers, inboxHandlerMiddleware)

	// setup driver adapters
	if err = handlers.RegisterIntegrationEventHandlers(stream, msgEvtHandlerMiddleware); err != nil {
		return err
	}

	replyMsgHandlers := am.NewReplyMessageHandler(reg,orchestrator)
	msgReplyHandlerMiddleware := am.RawMessageHandlerWithMiddleware(replyMsgHandlers, inboxHandlerMiddleware)
	if err = handlers.RegisterReplyHandlers(stream, msgReplyHandlerMiddleware); err != nil {
		return err
	}
	
	outboxProcessor := tm.NewOutboxProcessor(jsStream, pg.NewOutboxStore("cosec.outbox", mono.DB()))
	go func() {
		if err := outboxProcessor.Start(ctx); err != nil {
			logger := mono.Logger()
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
