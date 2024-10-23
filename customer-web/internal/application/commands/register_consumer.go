package commands

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
)

type RegisterConsumer struct {
	Name string
}

type RegisterConsumerHandler struct {
	consumers domain.ConsumerRepository
}

func NewRegisterConsumerHandler(consumers domain.ConsumerRepository) RegisterConsumerHandler {
	return RegisterConsumerHandler{
		consumers: consumers,
	}
}

func (h RegisterConsumerHandler) RegisterConsumer(ctx context.Context, cmd RegisterConsumer) (string, error) {
	consumerID, err := h.consumers.Register(ctx, domain.RegisterConsumer{
		Name: cmd.Name,
	})
	if err != nil {
		return "", err
	}

	return consumerID, nil
}
