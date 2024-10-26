package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderRepository struct {
	client orderpb.OrderServiceClient
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(conn *grpc.ClientConn) OrderRepository {
	return OrderRepository{client: orderpb.NewOrderServiceClient(conn)}
}

func (r OrderRepository) Create(ctx context.Context, create domain.CreateOrder) (string, error) {
	resp, err := r.client.CreateOrder(ctx, &orderpb.CreateOrderRequest{
		ConsumerID:   create.ConsumerID,
		RestaurantID: create.RestaurantID,
		DeliverAt:    timestamppb.New(create.DeliverAt),
		DeliverTo:    r.toAddressProto(create.DeliverTo),
		LineItems:    r.toMenuItemQuantitiesProto(create.LineItems),
	})
	if err != nil {
		return "", err
	}
	return resp.GetOrderID(), nil
}

func (r OrderRepository) Find(ctx context.Context, find domain.FindOrder) (*domain.Order, error) {
	orderResp, err := r.client.GetOrder(ctx, &orderpb.GetOrderRequest{
		OrderID: find.OrderID,
	})
	if err != nil {
		return nil, err
	}
	return r.toOrderDomain(orderResp.Order), nil
}

func (r OrderRepository) toOrderDomain(order *orderpb.Order) *domain.Order {
	return &domain.Order{
		OrderID:      order.OrderID,
		ConsumerID:   order.ConsumerID,
		RestaurantID: order.RestaurantID,
		Total:        int(order.OrderTotal),
		Status:       string(order.Status),
	}
}
func (r OrderRepository) toAddressProto(address domain.Address) *orderpb.Address {
	return &orderpb.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}

func (r OrderRepository) toMenuItemQuantitiesProto(quantities domain.MenuItemQuantities) *orderpb.MenuItemQuantities {
	lineItems := make(map[string]int64, len(quantities))
	for itemID, qty := range quantities {
		lineItems[itemID] = int64(qty)
	}
	return &orderpb.MenuItemQuantities{Items: lineItems}
}

func (r OrderRepository) toAddressDomain(address *consumerpb.Address) *domain.Address {
	return &domain.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}
