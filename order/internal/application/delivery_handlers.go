package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
)

type DliveryHandlers struct {
	ignoreUnimplementedDomainEvents
	deliveries domain.DeliveryRepository
}

func NewDliveryHandlers(deliveries domain.DeliveryRepository)*DliveryHandlers{
	return &DliveryHandlers{
		deliveries: deliveries,
	}
}

func(h DliveryHandlers)OnOrderCreated(ctx context.Context, event ddd.Event) error{
	orderCreated := event.(*domain.OrderCreated)
	return h.deliveries.CreateDelivery(ctx,domain.CreateDelivery{
		DeliveryID: orderCreated.Order.GetID(),
		RestaurantID: orderCreated.Order.RestaurantID,
		Address: orderCreated.Order.DeliverTo,
	})
}