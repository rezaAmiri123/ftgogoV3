package commands

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
	"github.com/stackus/errors"
)

type ConfirmCreateTicket struct {
	ID           string
}

type ConfirmCreateTicketHandler struct {
	tickets domain.TicketRepository
}

func NewConfirmCreateTicketHandler(tickets domain.TicketRepository) ConfirmCreateTicketHandler {
	return ConfirmCreateTicketHandler{
		tickets: tickets,
	}
}

func (h ConfirmCreateTicketHandler) ConfirmCreateTicket(ctx context.Context, cmd ConfirmCreateTicket) error {
	ticket, err := h.tickets.Find(ctx,cmd.ID)
	if err != nil {
		return err
	}

	err = ticket.ConfirmCreate()
	if err != nil {
		return err
	}

	return errors.Wrap(h.tickets.Update(ctx, ticket), "confirm create ticket query")
}
