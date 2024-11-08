package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/delivery/deliverypb"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/application/queries"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	app application.App
	deliverypb.UnimplementedDeliveryServiceServer
}

var _ deliverypb.DeliveryServiceServer = (*server)(nil)

func RegisterServer(app application.App, register grpc.ServiceRegistrar) error {
	deliverypb.RegisterDeliveryServiceServer(register, server{app: app})
	return nil
}

func (s server) CreateDelivery(ctx context.Context, request *deliverypb.CreateDeliveryRequest) (*deliverypb.CreateDeliveryResponse, error) {
	err := s.app.CreateDelivery(ctx, commands.CreateDelivery{
		ID:              request.GetID(),
		RestaurantID:    request.GetRestaurantID(),
		DeliveryAddress: s.toAddressDomain(request.GetDeliveryAddress()),
	})
	if err != nil {
		return nil, err
	}
	return &deliverypb.CreateDeliveryResponse{}, nil
}

func (s server) GetDelivery(ctx context.Context, request *deliverypb.GetDeliveryRequest) (*deliverypb.GetDeliveryResponse, error) {
	delivery, err := s.app.GetDelivery(ctx, queries.GetDelivery{ID: request.GetDeliveryID()})
	if err != nil {
		return nil, err
	}
	return &deliverypb.GetDeliveryResponse{
		Delivery: s.toDeliveryProto(delivery),
	}, nil
}

func (s server) ScheduleDelivery(ctx context.Context, request *deliverypb.ScheduleDeliveryRequest) (*deliverypb.ScheduleDeliveryResponse, error) {
	err := s.app.ScheduleDelivery(ctx, commands.ScheduleDelivery{
		ID:      request.GetID(),
		ReadyBy: request.ReadyBy.AsTime(),
	})
	if err != nil {
		return nil, err
	}
	return &deliverypb.ScheduleDeliveryResponse{}, nil
}

func (s server) SetCourierAvailability(ctx context.Context, request *deliverypb.SetCourierAvailabilityRequest) (*deliverypb.SetCourierAvailabilityResponse, error) {
	err := s.app.SetCourierAvailability(ctx, commands.SetCourierAvailability{
		CourierID: request.GetCourierID(),
		Available: request.GetAvailable(),
	})
	if err != nil {
		return nil, err
	}
	return &deliverypb.SetCourierAvailabilityResponse{
		Available: request.Available,
	}, nil
}

func (s server) toAddressDomain(address *deliverypb.Address) domain.Address {
	return domain.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}

func (s server) toAddressProto(address domain.Address) *deliverypb.Address {
	return &deliverypb.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}

func (s server) toDeliveryProto(delivery *domain.Delivery) *deliverypb.Delivery {
	return &deliverypb.Delivery{
		DeliveryID:        delivery.ID(),
		RestaurantID:      delivery.RestaurantID,
		AssignedCourierID: delivery.AssignedCourierID,
		Status:            delivery.Status.String(),
		PickUpAddress:     s.toAddressProto(delivery.PickUpAddress),
		DeliveryAddress:   s.toAddressProto(delivery.DeliveryAddress),
		PickupTime:        timestamppb.New(delivery.PickUpTime),
		ReadyBy:           timestamppb.New(delivery.ReadyBy),
	}
}
