package domain

import (
	"time"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/es"
	"github.com/stackus/errors"
)

const OrderAggregate = "order.OrderAggregate"

var (
	ErrOrderAlreadyCreated       = errors.Wrap(errors.ErrBadRequest, "the order cannot be recreated")
	ErrOrderIDCannotBeBlank      = errors.Wrap(errors.ErrBadRequest, "the order id cannot be blank")
	ErrConsumerIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the consumer id cannot be blank")
	ErrRestaurantIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the retaurant id cannot be blank")
	ErrLineItemsCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the Line items cannot be blank")
	ErrDeliverAtCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the deliver at cannot be blank")
	ErrDeliverToCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the deliverTo cannot be blank")
	ErrOrderInvalidStatus        = errors.Wrap(errors.ErrFailedPrecondition, "order status does not allow action")
)

type Order struct {
	es.Aggregate
	ConsumerID   string
	RestaurantID string
	TicketID     string
	LineItems    []LineItem
	Status       OrderStatus
	DeliverAt    time.Time
	DeliverTo    Address
}

var _ interface {
	es.EventApplier
	// es.Snapshotter
} = (*Order)(nil)

func (Order) Key() string { return OrderAggregate }

func NewOrder(id string) *Order {
	return &Order{
		Aggregate: es.NewAggregate(id, OrderAggregate),
	}
}

func (o *Order) CreateOrder(consumerID, restaurantID string, lineItems []LineItem, deliverAt time.Time, deliverTo Address) error {
	if o.Status != UnknownOrderStatus {
		return ErrOrderAlreadyCreated
	}

	if consumerID == "" {
		return ErrConsumerIDCannotBeBlank
	}
	if restaurantID == "" {
		return ErrRestaurantIDCannotBeBlank
	}
	if len(lineItems) == 0 {
		return ErrLineItemsCannotBeBlank
	}
	if deliverAt == (time.Time{}) {
		return ErrDeliverAtCannotBeBlank
	}
	if deliverTo == (Address{}) {
		return ErrDeliverToCannotBeBlank
	}
	o.AddEvent(OrderCreatedEvent, &OrderCreated{
		ConsumerID:   consumerID,
		RestaurantID: restaurantID,
		LineItems:    lineItems,
		DeliverAt:    deliverAt,
		DeliverTo:    deliverTo,
	})

	return nil
}

func (o *Order) OrderTotal() int {
	total := 0
	for _, item := range o.LineItems {
		total += item.GetTotal()
	}

	return total
}

func (o *Order) ApproveOrder(ticketID string) error {
	// if o.Status != ApprovalPending {
	// 	return ErrOrderInvalidStatus
	// }

	o.AddEvent(OrderApprovedEvent, &OrderApproved{
		TicketID: ticketID,
	})

	return nil
}

func (o *Order) ApplyEvent(event ddd.Event) error {
	switch payload := event.Payload().(type) {
	case *OrderCreated:
		o.ConsumerID = payload.ConsumerID
		o.RestaurantID = payload.RestaurantID
		o.LineItems = payload.LineItems
		o.DeliverAt = payload.DeliverAt
		o.DeliverTo = payload.DeliverTo
		o.Status = ApprovalPending
	case *OrderApproved:
		o.TicketID = payload.TicketID
		o.Status = Approved
	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", o, event.EventName(), payload)
	}
	return nil
}

func (o *Order) ApplySnapshot(snapshot es.Snapshot) error {
	switch ss := snapshot.(type) {
	case *OrderV1:
		o.ConsumerID = ss.ConsumerID
		o.RestaurantID = ss.RestaurantID
		o.TicketID = ss.TicketID
		o.LineItems = ss.LineItems
		o.Status = ss.Status
		o.DeliverAt = ss.DeliverAt
		o.DeliverTo = ss.DeliverTo
	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", o, snapshot)
	}
	return nil
}

func (o *Order) ToSnapshot() es.Snapshot {
	return &OrderV1{
		ConsumerID:   o.ConsumerID,
		RestaurantID: o.RestaurantID,
		TicketID:     o.TicketID,
		LineItems:    o.LineItems,
		Status:       o.Status,
		DeliverAt:    o.DeliverAt,
		DeliverTo:    o.DeliverTo,
	}
}
