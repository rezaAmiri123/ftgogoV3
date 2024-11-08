package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
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
	GetConsumerAddress struct {
		ConsumerID string
		AddressID  string
	}
	ValidateOrderByConsumer struct {
		ConsumerID string
		OrderID    string
		OrderTotal int
	}

	Application struct {
		consumers       domain.ConsumerRepository
		domainPublisher ddd.EventPublisher[ddd.AggregateEvent]
	}
)

//go:generate mockery --name App
type App interface {
	RegisterConsumer(ctx context.Context, register RegisterConsumer) error
	GetConsumer(ctx context.Context, get GetConsumer) (*domain.Consumer, error)
	UpdateConsumerAddress(ctx context.Context, update UpdateConsumerAddress) error
	RemoveConsumerAddress(ctx context.Context, remove RemoveConsumerAddress) error
	GetConsumerAddress(ctx context.Context, get GetConsumerAddress) (domain.Address, error)
	ValidateOrderByConsumer(ctx context.Context, validate ValidateOrderByConsumer) error
}

var _ App = (*Application)(nil)

func New(
	consumers domain.ConsumerRepository,
	domainPublisher ddd.EventPublisher[ddd.AggregateEvent],
) *Application {
	return &Application{
		consumers:       consumers,
		domainPublisher: domainPublisher,
	}
}

func (a Application) RegisterConsumer(ctx context.Context, register RegisterConsumer) error {
	consumer, err := domain.RegisterConsumer(register.ID, register.Name)
	if err != nil {
		return err
	}

	// publishe domain events
	err = a.consumers.Save(ctx, consumer)
	if err != nil {
		return err
	}

	return a.domainPublisher.Publish(ctx, consumer.Events()...)
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

func (a Application) RemoveConsumerAddress(ctx context.Context, remove RemoveConsumerAddress) error {
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

func (a Application) GetConsumerAddress(ctx context.Context, get GetConsumerAddress) (domain.Address, error) {
	consumer, err := a.consumers.Find(ctx, get.ConsumerID)
	if err != nil {
		return domain.Address{}, err
	}
	address, err := consumer.GetAddress(get.AddressID)
	if err != nil {
		return domain.Address{}, err
	}
	return address, nil
}

func (a Application) ValidateOrderByConsumer(ctx context.Context, validate ValidateOrderByConsumer) error {
	consumer, err := a.consumers.Find(ctx, validate.ConsumerID)
	if err != nil {
		return err
	}

	return consumer.ValidateOrderByConsumer(validate.OrderTotal)
}
