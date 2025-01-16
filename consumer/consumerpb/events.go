package consumerpb

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
)

const (
	ConsumerAggregateChannel = "ftgogo.consumers.events.Consumer"

	ConsumerRegisteredEvent = "consumerapi.ConsumerRegistered"

	CommandChannel = "ftgogo.consumers.commands"

	AuthorizeConsumerCommand = "consumerapi.AuthorizeConsumer"
)

func Registration(reg registry.Registry) (err error) {
	serde := serdes.NewProtoSerde(reg)

	// Consumer events
	if err = serde.Register(&ConsumerRegistred{}); err != nil {
		return err
	}

	// Commands
	if err = serde.Register(&AuthorizeCustomer{}); err != nil {
		return err
	}
	return nil
}

func (*ConsumerRegistred) Key() string { return ConsumerRegisteredEvent }

func (*AuthorizeCustomer) Key() string { return AuthorizeConsumerCommand }
