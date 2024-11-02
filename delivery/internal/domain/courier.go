package domain

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/stackus/errors"
)

var (
	ErrCourierNotFound = errors.Wrap(errors.ErrNotFound, "courier not found")
)

type Courier struct {
	ddd.AggregateBase
	Plan      Plan
	Available bool
}

func (c *Courier) AddAction(action Action) {
	c.Plan.Add(action)
}

func (c *Courier) CancelDelivery(deliveryID string) {
	c.Plan.RemoveDelivery(deliveryID)
}
