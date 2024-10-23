package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
	"google.golang.org/grpc"
)

type ConsumerRepository struct {
	client consumerpb.ConsumerServiceClient
}

var _ domain.ConsumerRepository = (*ConsumerRepository)(nil)

func NewConsumerRepository(conn *grpc.ClientConn) ConsumerRepository {
	return ConsumerRepository{client: consumerpb.NewConsumerServiceClient(conn)}
}

func (r ConsumerRepository) ValidateOrderByConsumer(ctx context.Context, validate domain.ValidateOrderByConsumer) error {
	_, err := r.client.ValidateOrderByConsumer(ctx, &consumerpb.ValidateOrderByConsumerRequest{
		ConsumerID: validate.ConsumerID,
		OrderID:    validate.OrderID,
		OrderTotal: int64(validate.OrderTotal),
	})
	return err
}
