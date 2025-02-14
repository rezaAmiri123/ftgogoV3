package handlers

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/errorsotel"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type domainHandlers[T ddd.AggregateEvent] struct {
	publisher am.EventPublisher
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*domainHandlers[ddd.AggregateEvent])(nil)

func NewDomainEventHandlers(publisher am.EventPublisher) ddd.EventHandler[ddd.AggregateEvent] {
	return &domainHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}
func RegisterDomainEventHandlers(eventHandlers ddd.EventHandler[ddd.AggregateEvent], subscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	subscriber.Subscribe(eventHandlers,
		kitchenpb.TicketAcceptedEvent,
	)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling domain event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled domain event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling domain event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))

	switch event.EventName() {
	case kitchenpb.TicketAcceptedEvent:
		return h.onTicketAccepted(ctx, event)
	}
	return nil
}

func (h domainHandlers[T]) onTicketAccepted(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.TicketAccepted)
	return h.publisher.Publish(ctx, kitchenpb.TicketAggregateChannel,
		ddd.NewEvent(kitchenpb.TicketAcceptedEvent, &kitchenpb.TicketAccepted{
			TicketID:   payload.Ticket.ID(),
			OrderID:    payload.Ticket.OrderID,
			AcceptedAt: timestamppb.New(payload.Ticket.AcceptedAt),
			ReadyBy:    timestamppb.New(payload.Ticket.ReadyBy),
		}),
	)
}
