package commands

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
	"github.com/stackus/errors"
)

type CreateTicket struct {
	ID           string
	OrderID      string
	RestaurantID string
	LineItems    []domain.LineItem
}

type CreateTicketHandler struct {
	tickets domain.TicketRepository
}

func NewCreateTicketHandler(tickets domain.TicketRepository) CreateTicketHandler {
	return CreateTicketHandler{
		tickets: tickets,
	}
}

func (h CreateTicketHandler) CreateTicket(ctx context.Context, cmd CreateTicket) error {
	ticket, err := domain.CreateTicket(cmd.ID, cmd.RestaurantID, cmd.OrderID, cmd.LineItems)
	if err != nil {
		return err
	}
	return errors.Wrap(h.tickets.Save(ctx, ticket), "create ticket query")
}
