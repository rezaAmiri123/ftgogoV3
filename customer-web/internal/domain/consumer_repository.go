package domain

import "context"

type (
	RegisterConsumer struct {
		Name string
	}

	FindConsumer struct {
		ConsumerID string
	}

	UpdateConsumerAddress struct {
		ConsumerID string
		AddressID  string
		Address    Address
	}
	FindConsumerAddress struct {
		ConsumerID string
		AddressID  string
	}

)

type ConsumerRepository interface {
	Register(ctx context.Context, register RegisterConsumer) (string, error)
	Find(ctx context.Context, find FindConsumer) (*Consumer, error)
	UpdateConsumerAddress(ctx context.Context, update UpdateConsumerAddress) error
	FindAddress(ctx context.Context, find FindConsumerAddress) (*Address, error)
}
