package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
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
		ID:        restaurant.GetRestaurantID(),
		Name:      restaurant.GetName(),
		MenuItems: r.toMenuItemsDomain(restaurant.MenuItems),
	}, nil

}

func (r RestaurantRepository) toMenuItemsDomain(lineItems []*restaurantpb.MenuItem) []domain.MenuItem {
	resp := make([]domain.MenuItem, len(lineItems))
	for i, lineItem := range lineItems {
		resp[i] = domain.MenuItem{
			ID:    lineItem.GetID(),
			Name:  lineItem.GetName(),
			Price: int(lineItem.GetPrice()),
		}
	}
	return resp
}
