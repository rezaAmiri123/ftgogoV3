package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
)

type DeliveryHandlers struct {
	ignoreUnimplementedDomainEvents
	deliveries domain.DeliveryRepository
}

func NewDeliveryHandlers(deliveries domain.DeliveryRepository) *DeliveryHandlers {
	return &DeliveryHandlers{
		deliveries: deliveries,
	}
}

func (h DeliveryHandlers) OnTicketAccepted(ctx context.Context, event ddd.Event) error {
	ticketAcceted := event.(*domain.TicketAccepted)
	return h.deliveries.ScheduleDelivery(ctx, domain.ScheduleDelivery{
		ID:      ticketAcceted.Ticket.ID,
		ReadyBy: ticketAcceted.Ticket.ReadyBy,
	})
}
