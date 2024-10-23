package domain

import (
	"context"
)

type CreateTicket struct{
	ID string
	RestaurantID string
	TicketDetail []LineItem
}

type KitchenRepository interface {
	CreateTicket(ctx context.Context, create CreateTicket)error
	ConfirmCreateTicket(ctx context.Context, ticketID string)error
}
