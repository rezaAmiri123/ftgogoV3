package commands

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
)

type ScheduleDelivery struct {
	ID      string
	ReadyBy time.Time
}

type ScheduleDeliveryHandler struct {
	deliveries domain.DeliveryRepository
	couriers   domain.CourierRepository
}

func NewScheduleDeliveryHandler(
	deliveries domain.DeliveryRepository,
	couriers domain.CourierRepository,
) ScheduleDeliveryHandler {
	return ScheduleDeliveryHandler{
		deliveries: deliveries,
		couriers:   couriers,
	}
}

func (h ScheduleDeliveryHandler) ScheduleDelivery(ctx context.Context, cmd ScheduleDelivery) error {
	delivery, err := h.deliveries.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	courier, err := h.couriers.FindFirstAvailable(ctx)
	if err != nil {
		return err
	}

	courier.Plan.Add(domain.Action{
		DeliveryID: delivery.ID,
		ActionType: domain.PickUp,
		Address:    delivery.PickUpAddress,
		When:       cmd.ReadyBy,
	})
	courier.Plan.Add(domain.Action{
		DeliveryID: delivery.ID,
		ActionType: domain.DropOff,
		Address:    delivery.DeliveryAddress,
		When:       cmd.ReadyBy.Add(30 * time.Minute),
	})

	err = h.couriers.Update(ctx, courier)
	if err != nil {
		return err
	}

	delivery.Schedule(cmd.ReadyBy, courier.ID)

	return h.deliveries.Update(ctx, delivery)
}
