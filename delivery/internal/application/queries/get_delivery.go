package queries

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
)

type GetDelivery struct {
	ID string
}

type GetDeliveryHandler struct {
	deliveries domain.DeliveryRepository
}

func NewGetDeliveryHandler(deliveries domain.DeliveryRepository) GetDeliveryHandler {
	return GetDeliveryHandler{
		deliveries: deliveries,
	}
}

func (h GetDeliveryHandler) GetDelivery(ctx context.Context, query GetDelivery) (*domain.Delivery, error) {
	return h.deliveries.Find(ctx, query.ID)
}
