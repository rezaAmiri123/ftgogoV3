package domain

import (
	"context"
)

type TicketRepository interface {
	Find(ctx context.Context, ticketID string) (*Ticket, error)
	Save(ctx context.Context, ticket *Ticket) error
	Update(ctx context.Context, ticket *Ticket) error
}
