package handlers

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type domainHandlers[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*domainHandlers[ddd.AggregateEvent])(nil)

func NewDomainEventHandlers(publisher am.MessagePublisher[ddd.Event]) *domainHandlers[ddd.AggregateEvent] {
	return &domainHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}
func RegisterDomainEventHandlers(eventHandlers ddd.EventHandler[ddd.AggregateEvent], subscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	subscriber.Subscribe(eventHandlers,
		kitchenpb.TicketAcceptedEvent,
	)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) error {
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
