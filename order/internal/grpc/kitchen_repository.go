package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
	"google.golang.org/grpc"
)

type KitchenRepository struct {
	client kitchenpb.KitchenServiceClient
}

var _ domain.KitchenRepository = (*KitchenRepository)(nil)

func NewKitchenRepository(conn *grpc.ClientConn) KitchenRepository {
	return KitchenRepository{client: kitchenpb.NewKitchenServiceClient(conn)}
}

func (r KitchenRepository) CreateTicket(ctx context.Context, create domain.CreateTicket) error {
	_, err := r.client.CreateTicket(ctx, &kitchenpb.CreateTicketRequest{
		OrderID:      create.OrderID,
		RestaurantID: create.RestaurantID,
		LineItems:    r.toLineItemsProto(create.TicketDetail),
	})
	return err
}

func (r KitchenRepository) ConfirmCreateTicket(ctx context.Context, ticketID string) error {
	_, err := r.client.ConfirmCreateTicket(ctx, &kitchenpb.ConfirmCreateTicketRequest{
		TicketID: ticketID,
	})
	return err
}

func (r KitchenRepository) toLineItemsProto(lineItems []domain.LineItem) []*kitchenpb.LineItem {
	resp := make([]*kitchenpb.LineItem, len(lineItems))
	for i, lineItem := range lineItems {
		resp[i] = &kitchenpb.LineItem{
			MenuItemID: lineItem.MenuItemID,
			Name:       lineItem.Name,
			Quantity:   int64(lineItem.Quantity),
		}
	}
	return resp
}
