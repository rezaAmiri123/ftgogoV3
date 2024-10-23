package domain

import (
	"time"

	"github.com/stackus/errors"
)

type Order struct {
	ID           string
	ConsumerID   string
	RestaurantID string
	TicketID     string
	LineItems    []LineItem
	Status       OrderStatus
	DeliverAt    time.Time
	DeliverTo    Address
}

var (
	ErrOrderIDCannotBeBlank      = errors.Wrap(errors.ErrBadRequest, "the order id cannot be blank")
	ErrConsumerIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the consumer id cannot be blank")
	ErrRestaurantIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the retaurant id cannot be blank")
	ErrLineItemsCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the Line items cannot be blank")
	ErrDeliverAtCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the deliver at cannot be blank")
	ErrDeliverToCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the deliverTo cannot be blank")
	ErrOrderInvalidStatus        = errors.Wrap(errors.ErrFailedPrecondition, "order status does not allow action")
)

func CreateOrder(id, consumerID, restaurantID string, lineItems []LineItem, deliverAt time.Time, deliverTo Address) (*Order, error) {
	if id == "" {
		return nil, ErrOrderIDCannotBeBlank
	}
	if consumerID == "" {
		return nil, ErrConsumerIDCannotBeBlank
	}
	if restaurantID == "" {
		return nil, ErrRestaurantIDCannotBeBlank
	}
	if len(lineItems) == 0 {
		return nil, ErrLineItemsCannotBeBlank
	}
	if deliverAt == (time.Time{}) {
		return nil, ErrDeliverAtCannotBeBlank
	}
	if deliverTo == (Address{}) {
		return nil, ErrDeliverToCannotBeBlank
	}

	order := &Order{
		ID:           id,
		ConsumerID:   consumerID,
		RestaurantID: restaurantID,
		LineItems:    lineItems,
		Status:       ApprovalPending,
		DeliverAt:    deliverAt,
		DeliverTo:    deliverTo,
	}

	return order, nil
}

func (o *Order) OrderTotal() int {
	total := 0
	for _, item := range o.LineItems {
		total += item.GetTotal()
	}

	return total
}

func (o *Order) ApproveOrder(ticketID string) error {
	if o.Status != ApprovalPending {
		return ErrOrderInvalidStatus
	}
	o.TicketID = ticketID
	o.Status = Approved
	return nil
}
