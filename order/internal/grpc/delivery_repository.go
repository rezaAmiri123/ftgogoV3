package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/delivery/deliverypb"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
	"google.golang.org/grpc"
)

type DeliveryRepository struct {
	client deliverypb.DeliveryServiceClient
}

var _ domain.DeliveryRepository = (*DeliveryRepository)(nil)

func NewDeliveryRepository(conn *grpc.ClientConn) DeliveryRepository {
	return DeliveryRepository{client: deliverypb.NewDeliveryServiceClient(conn)}
}

func (r DeliveryRepository) CreateDelivery(ctx context.Context, create domain.CreateDelivery) error {
	_, err := r.client.CreateDelivery(ctx, &deliverypb.CreateDeliveryRequest{
		ID:              create.DeliveryID,
		RestaurantID:    create.RestaurantID,
		DeliveryAddress: r.toAddressProto(create.Address),
	})
	return err
}

func (r DeliveryRepository) toAddressProto(address domain.Address) *deliverypb.Address {
	return &deliverypb.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}
