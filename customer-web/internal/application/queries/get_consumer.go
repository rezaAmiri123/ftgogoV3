package queries

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
)

type GetConsumer struct {
	ConsumerID string
}

type GetConsumerHandler struct {
	consumers domain.ConsumerRepository
}

func NewGetConsumerHandler(consumers domain.ConsumerRepository) GetConsumerHandler {
	return GetConsumerHandler{
		consumers: consumers,
	}
}

func (h GetConsumerHandler) GetConsumer(ctx context.Context, query GetConsumer) (*domain.Consumer, error) {
	consumer, err := h.consumers.Find(ctx, domain.FindConsumer{
		ConsumerID: query.ConsumerID,
	})
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
