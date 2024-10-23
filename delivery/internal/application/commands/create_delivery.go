package commands

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
)

type CreateDelivery struct {
	ID              string
	RestaurantID    string
	DeliveryAddress domain.Address
}

type CreateDeliveryHandler struct {
	deliveries  domain.DeliveryRepository
	restauants domain.RestaurantRepository
}

func NewCreateDeliveryHandler(
	deliveries domain.DeliveryRepository,
	restauants domain.RestaurantRepository,
) CreateDeliveryHandler {
	return CreateDeliveryHandler{
		deliveries:  deliveries,
		restauants: restauants,
	}
}

func(h CreateDeliveryHandler)CreateDelivery(ctx context.Context, cmd CreateDelivery)error{
	restaurant,err := h.restauants.Find(ctx,cmd.RestaurantID)
	if err != nil{
		return err
	}

	delivery, err := domain.CreateDelivery(cmd.ID,restaurant.RestaurantID,restaurant.Address,cmd.DeliveryAddress)
	if err != nil{
		return err
	}
	return h.deliveries.Save(ctx,delivery)
}
