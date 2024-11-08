package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/restaurant/restaurantpb"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	restaurantpb.UnimplementedRestaurantServiceServer
}

var _ restaurantpb.RestaurantServiceServer = (*server)(nil)

func RegisterServer(app application.App, register grpc.ServiceRegistrar) error {
	restaurantpb.RegisterRestaurantServiceServer(register, server{app: app})
	return nil
}

func (s server) CreateRestaurant(ctx context.Context, request *restaurantpb.CreateRestaurantRequest) (*restaurantpb.CreateRestaurantResponse, error) {
	id := uuid.New().String()

	err := s.app.CreateRestaurant(ctx, application.CreateRestaurant{
		ID:      id,
		Name:    request.GetName(),
		Address: s.toAddressDomain(request.GetAddress()),
	})
	if err != nil {
		return nil, err
	}
	return &restaurantpb.CreateRestaurantResponse{RestaurantID: id}, nil
}

func (s server) GetRestaurant(ctx context.Context, request *restaurantpb.GetRestaurantRequest) (*restaurantpb.GetRestaurantResponse, error) {
	restaurant, err := s.app.GetRestaurant(ctx, application.GetRestaurant{
		ID: request.GetRestaurantID(),
	})
	if err != nil {
		return nil, err
	}
	return &restaurantpb.GetRestaurantResponse{
		RestaurantID: restaurant.ID(),
		Name:         restaurant.Name,
		Address:      s.toAddressProto(restaurant.Address),
		MenuItems:    s.toMenuItemsProto(restaurant.MenuItems),
	}, nil
}

func (s server) UpdateMenuItem(ctx context.Context, request *restaurantpb.UpdateMenuItemRequest) (*restaurantpb.UpdateMenuItemResponse, error) {
	err := s.app.UpdateMenuItem(ctx, application.UpdateMenuItem{
		RestaurantID: request.GetRestaurantID(),
		MenuItems:    s.toMenuItemsDomain(request.GetMenuItem()),
	})
	if err != nil {
		return nil, err
	}
	return &restaurantpb.UpdateMenuItemResponse{}, nil
}

func (s server) toAddressProto(address domain.Address) *restaurantpb.Address {
	return &restaurantpb.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}

func (s server) toAddressDomain(address *restaurantpb.Address) domain.Address {
	return domain.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}

func (s server) toMenuItemsProto(menuItems map[string]domain.MenuItem) []*restaurantpb.MenuItem {
	menuItemsProto := make([]*restaurantpb.MenuItem, 0, len(menuItems))
	for key, menuItem := range menuItems {
		menuItemsProto = append(menuItemsProto, &restaurantpb.MenuItem{
			ID:    key,
			Name:  menuItem.Name,
			Price: int64(menuItem.Price),
		})
	}
	return menuItemsProto
}

func (s server) toMenuItemsDomain(menuItems []*restaurantpb.MenuItem) []domain.MenuItem {
	menuItemsDomain := make([]domain.MenuItem, 0, len(menuItems))
	for _, menuItem := range menuItems {
		menuItemsDomain = append(menuItemsDomain, domain.MenuItem{
			ID:    menuItem.GetID(),
			Name:  menuItem.GetName(),
			Price: int(menuItem.GetPrice()),
		})
	}
	return menuItemsDomain
}
