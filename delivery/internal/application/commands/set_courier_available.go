package commands

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
)

type SetCourierAvailability struct {
	CourierID string
	Available bool
}

type SetCourierAvailabilityHandler struct {
	couriers domain.CourierRepository
}

func NewSetCourierAvailabilityHandler(couriers domain.CourierRepository) SetCourierAvailabilityHandler {
	return SetCourierAvailabilityHandler{
		couriers: couriers,
	}
}

func (h SetCourierAvailabilityHandler) SetCourierAvailability(ctx context.Context, cmd SetCourierAvailability) error {
	courier, err := h.couriers.FindOrCreate(ctx, cmd.CourierID)
	if err != nil {
		return err
	}
	
	courier.Available = cmd.Available
	return h.couriers.Update(ctx, courier)
}
