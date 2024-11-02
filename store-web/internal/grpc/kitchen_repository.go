package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/domain"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type KitchenRepository struct {
	client kitchenpb.KitchenServiceClient
}

var _ domain.KitchenRepository = (*KitchenRepository)(nil)

func NewKitchenRepository(conn *grpc.ClientConn) KitchenRepository {
	return KitchenRepository{client: kitchenpb.NewKitchenServiceClient(conn)}
}

func (r KitchenRepository) AcceptTicket(ctx context.Context, accept domain.AcceptTicket) error {
	_, err := r.client.AcceptTicket(ctx, &kitchenpb.AcceptTicketRequest{
		TicketID: accept.ID,
		ReadyBy:  timestamppb.New(accept.ReadyBy),
	})
	return err
}
