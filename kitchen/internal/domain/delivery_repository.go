package domain

import (
	"context"
	"time"
)

type ScheduleDelivery struct {
	ID      string
	ReadyBy time.Time
}

type DeliveryRepository interface {
	ScheduleDelivery(ctx context.Context, schedule ScheduleDelivery) error
}
