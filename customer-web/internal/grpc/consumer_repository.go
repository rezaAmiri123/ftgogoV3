package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
	"google.golang.org/grpc"
)

type ConsumerRepository struct {
	client consumerpb.ConsumerServiceClient
}

var _ domain.ConsumerRepository = (*ConsumerRepository)(nil)

func NewConsumerRepository(conn *grpc.ClientConn) ConsumerRepository {
	return ConsumerRepository{client: consumerpb.NewConsumerServiceClient(conn)}
}

func (r ConsumerRepository) Register(ctx context.Context, register domain.RegisterConsumer) (string, error) {
	resp, err := r.client.RegisterConsumer(ctx, &consumerpb.RegisterConsumerRequest{
		Name: register.Name,
	})
	if err != nil {
		return "", err
	}
	return resp.GetConsumerID(), nil
}
