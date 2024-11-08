package domain

import "time"

const (
	OrderCreatedEvent  = "order.OrderCreated"
	OrderApprovedEvent = "order.OrderApproved"
)

type OrderCreated struct {
	ConsumerID   string
	RestaurantID string
	LineItems    []LineItem
	DeliverAt    time.Time
	DeliverTo    Address
}

func (OrderCreated) Key() string { return OrderCreatedEvent }

type OrderApproved struct {
	TicketID string
}

func (OrderApproved) Key() string { return OrderApprovedEvent }
