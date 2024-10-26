package commands

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
)

type AddConsumerAddress struct {
	ConsumerID string
	AddressID  string
	Address    domain.Address
}

type AddConsumerAddressHandler struct {
	consumers domain.ConsumerRepository
}

func NewAddConsumerAddressHandler(consumers domain.ConsumerRepository) AddConsumerAddressHandler {
	return AddConsumerAddressHandler{
		consumers: consumers,
	}
}

func (h AddConsumerAddressHandler) AddConsumerAddress(ctx context.Context, cmd AddConsumerAddress) error {
	err := h.consumers.UpdateConsumerAddress(ctx, domain.UpdateConsumerAddress{
		ConsumerID: cmd.ConsumerID,
		AddressID:  cmd.AddressID,
		Address:    cmd.Address,
	})
	return err
}
