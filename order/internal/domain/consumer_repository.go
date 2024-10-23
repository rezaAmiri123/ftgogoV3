package domain

import (
	"context"
)

type ValidateOrderByConsumer struct {
	ConsumerID string
	OrderID    string
	OrderTotal int
}

type ConsumerRepository interface {
	ValidateOrderByConsumer(ctx context.Context, validate ValidateOrderByConsumer) error
}
