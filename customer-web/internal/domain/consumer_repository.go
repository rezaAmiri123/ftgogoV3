package domain

import "context"

type RegisterConsumer struct {
	Name string
}

type ConsumerRepository interface {
	Register(ctx context.Context, register RegisterConsumer) (string, error)
}