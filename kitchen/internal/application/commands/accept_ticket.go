package commands

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
)

type AcceptTicket struct {
	ID      string
	ReadyBy time.Time
}

type AcceptTicketHandler struct {
	tickets         domain.TicketRepository
	domainPublisher ddd.EventPublisher
}

func NewAcceptTicketHandler(tickets domain.TicketRepository, domainPublisher ddd.EventPublisher) AcceptTicketHandler {
	return AcceptTicketHandler{
		tickets:         tickets,
		domainPublisher: domainPublisher,
	}
}

func (h AcceptTicketHandler) AcceptTicket(ctx context.Context, cmd AcceptTicket) error {
	ticket, err := h.tickets.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = ticket.Accpept(cmd.ReadyBy)
	if err != nil {
		return err
	}

	err = h.tickets.Update(ctx, ticket)
	if err != nil {
		return err
	}

	return h.domainPublisher.Publish(ctx, ticket.GetEvents()...)
}
