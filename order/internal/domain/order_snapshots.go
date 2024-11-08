package domain

import "time"

type OrderV1 struct {
	ConsumerID   string
	RestaurantID string
	TicketID     string
	LineItems    []LineItem
	Status       OrderStatus
	DeliverAt    time.Time
	DeliverTo    Address
}

func (OrderV1) SnapshotName() string { return "order.OrderV1" }
