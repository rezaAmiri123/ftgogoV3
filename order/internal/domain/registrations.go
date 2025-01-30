package domain

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/es"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
)
func Registerations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Order
	err = serde.Register( Order{}, func(v any) error {
		order := v.(* Order)
		order.Aggregate = es.NewAggregate("", OrderAggregate)
		order.Status =  UnknownOrderStatus
		return nil
	})
	if err != nil {
		return err
	}

	// order events
	if err = serde.Register(OrderCreated{}); err != nil {
		return err
	}
	if err = serde.Register(OrderApproved{}); err != nil {
		return err
	}

	return nil
}
