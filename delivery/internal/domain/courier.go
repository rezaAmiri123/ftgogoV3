package domain

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/stackus/errors"
)

const CourierAggregate = "delivery.CoutierAggregate"

var (
	ErrCourierNotFound = errors.Wrap(errors.ErrNotFound, "courier not found")
)

type Courier struct {
	ddd.Aggregate
	Plan      Plan
	Available bool
}

func (Courier) Key() string { return CourierAggregate }

func NewCourier(id string)*Courier{
	return &Courier{
		Aggregate: ddd.NewAggregate(id,CourierAggregate),
	}
}

func (c *Courier) AddAction(action Action) {
	c.Plan.Add(action)
}

func (c *Courier) CancelDelivery(deliveryID string) {
	c.Plan.RemoveDelivery(deliveryID)
}
