package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
)

type (
	RegisterConsumer struct {
		ID   string
		Name string
	}
	GetConsumer struct {
		ID string
	}
	UpdateConsumerAddress struct {
		ConsumerID string
		AddressID  string
		Address    domain.Address
	}
	RemoveConsumerAddress struct {
		ConsumerID string
		AddressID  string
	}

	App interface {
		RegisterConsumer(ctx context.Context, register RegisterConsumer) error
		GetConsumer(ctx context.Context, get GetConsumer) (*domain.Consumer, error)
		UpdateConsumerAddress(ctx context.Context, update UpdateConsumerAddress) error
		RemoveConsumerAddress(ctx context.Context, remove RemoveConsumerAddress) error
	}

	Application struct {
		consumers domain.ConsumerRepository
	}
)

var _ App = (*Application)(nil)

func New(consumers domain.ConsumerRepository) *Application {
	return &Application{
		consumers: consumers,
	}
}

func (a Application) RegisterConsumer(ctx context.Context, register RegisterConsumer) error {
	consumer, err := domain.RegisterConsumer(register.ID, register.Name)
	if err != nil {
		return err
	}

	return a.consumers.Save(ctx, consumer)
}

func (a Application) GetConsumer(ctx context.Context, get GetConsumer) (*domain.Consumer, error) {
	return a.consumers.Find(ctx, get.ID)
}

func (a Application) UpdateConsumerAddress(ctx context.Context, update UpdateConsumerAddress) error {
	consumer, err := a.consumers.Find(ctx, update.ConsumerID)
	if err != nil {
		return err
	}
	err = consumer.UpdateAddress(update.AddressID, update.Address)
	if err != nil {
		return err
	}

	return a.consumers.Update(ctx, consumer)
}

func (a Application) RemoveConsumerAddress(ctx context.Context, remove RemoveConsumerAddress) error{
	consumer, err := a.consumers.Find(ctx, remove.ConsumerID)
	if err != nil {
		return err
	}
	err = consumer.RemoveAddress(remove.AddressID)
	if err != nil {
		return err
	}

	return a.consumers.Update(ctx, consumer)
}
