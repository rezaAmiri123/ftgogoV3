package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
)

type DeliveryHandlers[T ddd.AggregateEvent] struct {
	deliveries domain.DeliveryRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*DeliveryHandlers[ddd.AggregateEvent])(nil)

func NewDeliveryHandlers(deliveries domain.DeliveryRepository) DeliveryHandlers[ddd.AggregateEvent] {
	return DeliveryHandlers[ddd.AggregateEvent]{
		deliveries: deliveries,
	}
}

func (h DeliveryHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.TicketAcceptedEvent:
		return h.onTicketAccepted(ctx, event)
	}
	return nil
}

func (h DeliveryHandlers[T]) onTicketAccepted(ctx context.Context, event ddd.AggregateEvent) error {
	ticketAcceted := event.Payload().(*domain.TicketAccepted)
	return h.deliveries.ScheduleDelivery(ctx, domain.ScheduleDelivery{
		ID:      ticketAcceted.Ticket.ID(),
		ReadyBy: ticketAcceted.Ticket.ReadyBy,
	})
}
