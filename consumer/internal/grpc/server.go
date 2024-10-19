package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	consumerpb.UnimplementedConsumerServiceServer
}

var _ consumerpb.ConsumerServiceServer = (*server)(nil)

func RegisterServer(app application.App, register grpc.ServiceRegistrar) error {
	consumerpb.RegisterConsumerServiceServer(register, server{app: app})
	return nil
}
func (s server) RegisterConsumer(ctx context.Context, request *consumerpb.RegisterConsumerRequest) (*consumerpb.RegisterConsumerResponse, error) {
	id := uuid.New().String()

	err := s.app.RegisterConsumer(ctx, application.RegisterConsumer{
		ID:   id,
		Name: request.GetName(),
	})
	if err != nil {
		return nil, err
	}
	return &consumerpb.RegisterConsumerResponse{ConsumerID: id}, nil
}

func (s server) GetConsumer(ctx context.Context, request *consumerpb.GetConsumerRequest) (*consumerpb.GetConsumerResponse, error) {
	consumer, err := s.app.GetConsumer(ctx, application.GetConsumer{
		ID: request.GetConsumerID(),
	})
	if err != nil {
		return nil, err
	}
	return &consumerpb.GetConsumerResponse{
		ConsumerID: consumer.ID,
		Name:       consumer.Name,
	}, nil
}

func (s server) GetAddress(ctx context.Context, request *consumerpb.GetAddressRequest) (*consumerpb.GetAddressResponse, error) {
	address, err := s.app.GetConsumerAddress(ctx, application.GetConsumerAddress{
		ConsumerID: request.GetConsumerID(),
		AddressID:  request.GetAddressID(),
	})
	if err != nil {
		return nil, err
	}
	return &consumerpb.GetAddressResponse{
		ConsumerID: request.GetConsumerID(),
		AddressID:  request.GetAddressID(),
		Address: s.toAddressProto(address),
	}, nil
}

func (s server) UpdateAddress(ctx context.Context, request *consumerpb.UpdateAddressRequest) (*consumerpb.UpdateAddressResponse, error){
	err := s.app.UpdateConsumerAddress(ctx,application.UpdateConsumerAddress{
		ConsumerID: request.GetConsumerID(),
		AddressID: request.GetAddressID(),
		Address: s.toAddressDomain(request.GetAddress()),
	})
	if err != nil{
		return nil,err
	}
	return &consumerpb.UpdateAddressResponse{},nil
}

func (s server) RemoveAddress(ctx context.Context, request *consumerpb.RemoveAddressRequest) (*consumerpb.RemoveAddressResponse, error){
	err := s.app.RemoveConsumerAddress(ctx,application.RemoveConsumerAddress{
		ConsumerID: request.GetConsumerID(),
		AddressID: request.GetAddressID(),
	})
	if err != nil{
		return nil,err
	}
	return &consumerpb.RemoveAddressResponse{},nil
}

func (s server) toAddressProto(address domain.Address) *consumerpb.Address {
	return &consumerpb.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}

func (s server) toAddressDomain(address *consumerpb.Address) domain.Address {
	return domain.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}
