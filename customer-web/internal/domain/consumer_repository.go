package domain

import "context"

type (
	RegisterConsumer struct {
		Name string
	}

	FindConsumer struct {
		ConsumerID string
	}
)

type ConsumerRepository interface {
	Register(ctx context.Context, register RegisterConsumer) (string, error)
	Find(ctx context.Context, find FindConsumer) (*Consumer, error)
}
