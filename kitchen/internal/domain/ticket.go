package domain

import (
	"time"

	"github.com/stackus/errors"
)

var (
	ErrTicketIDCannotBeBlank     = errors.Wrap(errors.ErrBadRequest, "the ticket id cannot be blank")
	ErrRestaurantIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the restaurant id cannot be blank")
	ErrLineItemsCannotBeEmpty    = errors.Wrap(errors.ErrBadRequest, "the line items cannot be empty")
)

type Ticket struct {
	ID               string
	RestaurantID     string
	LineItems        []LineItem
	ReadyBy          time.Time
	AcceptedAt       time.Time
	PreparingTime    time.Time
	ReadyForPickUpAt time.Time
	PickedUpAt       time.Time
	Status           TicketStatus
	PerviousStatus   TicketStatus
}

func CreateTicket(id, restaurantID string, lineItems []LineItem) (*Ticket, error) {
	if id == "" {
		return nil, ErrTicketIDCannotBeBlank
	}
	if restaurantID == "" {
		return nil, ErrRestaurantIDCannotBeBlank
	}
	if len(lineItems) == 0 {
		return nil, ErrLineItemsCannotBeEmpty
	}
	return &Ticket{
		ID:           id,
		RestaurantID: restaurantID,
		LineItems:    lineItems,
		Status:       CreatePending,
	}, nil
}
