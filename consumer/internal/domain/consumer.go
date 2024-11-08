package domain

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/stackus/errors"
)

const ConsumerAggregate = "consumer.Consumer"

var (
	ErrConsumerIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the consumer id cannot be blank")
	ErrConsumerNameCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the consumer name cannot be blank")
	ErrAddressDoesNotExist       = errors.Wrap(errors.ErrBadRequest, "the address does not exists")
)

type Consumer struct {
	ddd.Aggregate
	Name      string
	Addresses map[string]Address
}

func (Consumer) Key() string { return ConsumerAggregate }

func NewConsumer(id string) *Consumer {
	return &Consumer{
		Aggregate: ddd.NewAggregate(id, ConsumerAggregate),
	}
}

func RegisterConsumer(id, name string) (*Consumer, error) {
	if id == "" {
		return nil, ErrConsumerIDCannotBeBlank
	}
	if name == "" {
		return nil, ErrConsumerNameCannotBeBlank
	}

	consumer := NewConsumer(id)
	consumer.Name = name
	consumer.Addresses = make(map[string]Address)

	consumer.AddEvent(ConsumerRegisteredEvent, &ConsumerRegistered{
		Consumer: consumer,
	})

	return consumer, nil
}

func (c *Consumer) UpdateAddress(id string, address Address) error {
	c.Addresses[id] = address
	return nil
}

func (c *Consumer) RemoveAddress(id string) error {
	if _, ok := c.Addresses[id]; !ok {
		return ErrAddressDoesNotExist
	}
	delete(c.Addresses, id)
	return nil
}

func (c *Consumer) GetAddress(id string) (Address, error) {
	address, ok := c.Addresses[id]
	if !ok {
		return Address{}, ErrAddressDoesNotExist
	}

	return address, nil
}

// ValidateOrderByConsumer domain method
func (c *Consumer) ValidateOrderByConsumer(orderTotal int) error {
	// ftgo: implement some business logic
	return nil
}
