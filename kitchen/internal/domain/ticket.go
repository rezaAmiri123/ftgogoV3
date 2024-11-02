package domain

import (
	"time"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/stackus/errors"
)

var (
	ErrTicketIDCannotBeBlank     = errors.Wrap(errors.ErrBadRequest, "the ticket id cannot be blank")
	ErrRestaurantIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the restaurant id cannot be blank")
	ErrLineItemsCannotBeEmpty    = errors.Wrap(errors.ErrBadRequest, "the line items cannot be empty")
	ErrTicketInvalidState        = errors.Wrap(errors.ErrFailedPrecondition, "ticket status does not allow action")
	ErrTicketReadyByBeforeNow    = errors.Wrap(errors.ErrInvalidArgument, "readyBy is in the past")
)

type Ticket struct {
	ddd.AggregateBase
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
	ticket := &Ticket{
		AggregateBase: ddd.AggregateBase{ID: id},
		RestaurantID:  restaurantID,
		LineItems:     lineItems,
		Status:        CreatePending,
	}

	return ticket, nil
}

func (t *Ticket) ConfirmCreate() error {
	if t.Status != CreatePending {
		return ErrTicketInvalidState
	}
	t.Status = AwaitingAcceptance
	return nil
}

func (t *Ticket) Accpept(readyBy time.Time) error {
	if t.Status != AwaitingAcceptance {
		return ErrTicketInvalidState
	}
	if time.Now().After(readyBy) {
		return ErrTicketReadyByBeforeNow
	}
	t.AcceptedAt = time.Now()
	t.ReadyBy = readyBy
	t.Status = Accepted // assume that this is the case; doesn't appear to be ever set in ftgo-kitchen-service

	t.AddEvent(&TicketAccepted{
		Ticket: t,
	})
	
	return nil
}
