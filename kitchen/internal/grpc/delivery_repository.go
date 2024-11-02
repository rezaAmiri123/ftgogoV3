package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/delivery/deliverypb"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DeliveryRepository struct {
	client deliverypb.DeliveryServiceClient
}

var _ domain.DeliveryRepository = (*DeliveryRepository)(nil)

func NewDeliveryRepository(conn *grpc.ClientConn) DeliveryRepository {
	return DeliveryRepository{client: deliverypb.NewDeliveryServiceClient(conn)}
}

func (r DeliveryRepository) ScheduleDelivery(ctx context.Context, schedule domain.ScheduleDelivery) error {
	_, err := r.client.ScheduleDelivery(ctx, &deliverypb.ScheduleDeliveryRequest{
		ID:      schedule.ID,
		ReadyBy: timestamppb.New(schedule.ReadyBy),
	})
	return err
}
