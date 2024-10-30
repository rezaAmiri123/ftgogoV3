package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/restaurant/restaurantpb"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/domain"
	"google.golang.org/grpc"
)

type RestaurantRepository struct {
	client restaurantpb.RestaurantServiceClient
}

var _ domain.RestaurantRepository = (*RestaurantRepository)(nil)

func NewRestaurantRepository(conn *grpc.ClientConn) RestaurantRepository {
	return RestaurantRepository{client: restaurantpb.NewRestaurantServiceClient(conn)}
}

func (r RestaurantRepository) Create(ctx context.Context, create domain.CreateRestaurant) (string, error) {
	resp, err := r.client.CreateRestaurant(ctx, &restaurantpb.CreateRestaurantRequest{
		Name:    create.Name,
		Address: r.toAddressProto(create.Address),
	})
	if err != nil {
		return "", err
	}
	return resp.GetRestaurantID(), nil
}

func (r RestaurantRepository) UpdateMenu(ctx context.Context, create domain.UpdateMenu) error {
	_, err := r.client.UpdateMenuItem(ctx, &restaurantpb.UpdateMenuItemRequest{
		RestaurantID: create.RestaurantID,
		MenuItem:     r.toMenuItemsProto(create.MenuItems),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r RestaurantRepository) toAddressProto(address domain.Address) *restaurantpb.Address {
	return &restaurantpb.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}

func (r RestaurantRepository) toMenuItemsProto(request []domain.MenuItem) []*restaurantpb.MenuItem {
	menuItems := make([]*restaurantpb.MenuItem, len(request))
	for i, menuItem := range request {
		menuItems[i] = &restaurantpb.MenuItem{
			ID:    menuItem.ID,
			Name:  menuItem.Name,
			Price: int64(menuItem.Price),
		}
	}
	return menuItems
}
