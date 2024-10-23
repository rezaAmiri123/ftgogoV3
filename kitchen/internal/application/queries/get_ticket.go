package queries

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
	"github.com/stackus/errors"
)

type GetTicket struct{
	ID string
}

type GetTicketHandler struct{
	tickets domain.TicketRepository
}

func NewGetTicketHandler(tickets domain.TicketRepository)GetTicketHandler{
	return GetTicketHandler{
		tickets: tickets,
	}
}

func(h GetTicketHandler)GetTicket(ctx context.Context, query GetTicket)(*domain.Ticket,error){
	ticket,err := h.tickets.Find(ctx,query.ID)
	return ticket,errors.Wrap(err, "get ticket query")
}