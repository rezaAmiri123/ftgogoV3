package orderpb

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
)

const (
	OrderAggregateChannel = "ftgogo.orders.events.Order"

	OrderCreatedEvent = "orderpb.OrderCreated"

	CommandChannel = "ftgogo.orders.commands"

	ApproveOrderCommand = "orders.ApproveOrderCommand"
	RejectOrderCommand  = "orders.RejectOrderCommand"
)

func Registeration(reg registry.Registry) (err error) {
	serde := serdes.NewProtoSerde(reg)

	// Order Events
	if err = serde.Register(&OrderCreated{}); err != nil {
		return err
	}

	// Commands
	if err = serde.Register(&ApproveOrder{}); err != nil {
		return err
	}
	if err = serde.Register(&RejectOrder{}); err != nil {
		return err
	}

	return nil
}

func (*OrderCreated) Key() string { return OrderCreatedEvent }

func (*ApproveOrder) Key() string { return ApproveOrderCommand }
func (*RejectOrder) Key() string  { return RejectOrderCommand }
