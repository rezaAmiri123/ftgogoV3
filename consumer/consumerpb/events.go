package consumerpb

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
)

const (
	ConsumerAggregateChannel = "ftgogo.consumers.events.Consumer"

	ConsumerRegisteredEvent = "consumerapi.ConsumerRegistered"
)

func Registration(reg registry.Registry) (err error) {
	serde := serdes.NewProtoSerde(reg)

	// Consumer events
	if err = serde.Register(&ConsumerRegistred{}); err != nil {
		return err
	}
	return nil
}

func (*ConsumerRegistred) Key() string { return ConsumerRegisteredEvent }
