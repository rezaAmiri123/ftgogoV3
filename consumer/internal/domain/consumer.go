package domain

import "github.com/stackus/errors"

var (
	ErrConsumerIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the consumer id cannot be blank")
	ErrConsumerNameCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the consumer name cannot be blank")
	ErrAddressDoesNotExist       = errors.Wrap(errors.ErrBadRequest, "the address does not exists")
)

type Consumer struct {
	ID        string
	Name      string
	Addresses map[string]Address
}

func RegisterConsumer(id, name string) (*Consumer, error) {
	if id == "" {
		return nil, ErrConsumerIDCannotBeBlank
	}
	if name == "" {
		return nil, ErrConsumerNameCannotBeBlank
	}

	return &Consumer{
		ID:        id,
		Name:      name,
		Addresses: make(map[string]Address),
	}, nil
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
