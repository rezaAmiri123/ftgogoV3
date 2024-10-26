package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
	"google.golang.org/grpc"
)

type ConsumerRepository struct {
	client consumerpb.ConsumerServiceClient
}

var _ domain.ConsumerRepository = (*ConsumerRepository)(nil)

func NewConsumerRepository(conn *grpc.ClientConn) ConsumerRepository {
	return ConsumerRepository{client: consumerpb.NewConsumerServiceClient(conn)}
}

func (r ConsumerRepository) Register(ctx context.Context, register domain.RegisterConsumer) (string, error) {
	resp, err := r.client.RegisterConsumer(ctx, &consumerpb.RegisterConsumerRequest{
		Name: register.Name,
	})
	if err != nil {
		return "", err
	}
	return resp.GetConsumerID(), nil
}

func (r ConsumerRepository) Find(ctx context.Context, find domain.FindConsumer) (*domain.Consumer, error) {
	consumer, err := r.client.GetConsumer(ctx, &consumerpb.GetConsumerRequest{
		ConsumerID: find.ConsumerID,
	})
	if err != nil {
		return nil, err
	}
	return &domain.Consumer{
		ConsumerID: consumer.GetConsumerID(),
		Name:       consumer.GetName(),
	}, nil
}

func (r ConsumerRepository) UpdateConsumerAddress(ctx context.Context, update domain.UpdateConsumerAddress) error {
	_, err := r.client.UpdateAddress(ctx, &consumerpb.UpdateAddressRequest{
		ConsumerID: update.ConsumerID,
		AddressID:  update.AddressID,
		Address:    r.toAddressProto(update.Address),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r ConsumerRepository) FindAddress(ctx context.Context, find domain.FindConsumerAddress) (*domain.Address, error) {
	address, err := r.client.GetAddress(ctx, &consumerpb.GetAddressRequest{
		ConsumerID: find.ConsumerID,
		AddressID:  find.AddressID,
	})
	if err != nil {
		return nil, err
	}
	return r.toAddressDomain(address.GetAddress()), nil
}

func (r ConsumerRepository) toAddressProto(address domain.Address) *consumerpb.Address {
	return &consumerpb.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}

func (r ConsumerRepository) toAddressDomain(address *consumerpb.Address) *domain.Address {
	return &domain.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}
