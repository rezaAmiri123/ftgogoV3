package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/restaurantpb"
	"google.golang.org/grpc"
)

type RestaurantRepository struct {
	client restaurantpb.RestaurantServiceClient
}

var _ domain.RestaurantRepository = (*RestaurantRepository)(nil)

func NewRestaurantRepository(conn *grpc.ClientConn) RestaurantRepository {
	return RestaurantRepository{client: restaurantpb.NewRestaurantServiceClient(conn)}
}

func (r RestaurantRepository) Find(ctx context.Context, restaurantID string) (*domain.Restaurant, error) {
	restaurant, err := r.client.GetRestaurant(ctx, &restaurantpb.GetRestaurantRequest{
		RestaurantID: restaurantID,
	})
	if err != nil {
		return nil, err
	}
	return &domain.Restaurant{
		RestaurantID: restaurant.GetRestaurantID(),
		Name:         restaurant.GetName(),
		Address:      r.toAddressDomain(restaurant.GetAddress()),
	}, nil

}

func (r RestaurantRepository) toAddressDomain(address *restaurantpb.Address) domain.Address {
	return domain.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}
