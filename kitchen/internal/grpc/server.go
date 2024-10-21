package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/application/queries"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	kitchenpb.UnimplementedKitchenServiceServer
}

var _ kitchenpb.KitchenServiceServer = (*server)(nil)

func RegisterServer(app application.App, register grpc.ServiceRegistrar) error {
	kitchenpb.RegisterKitchenServiceServer(register, server{app: app})
	return nil
}

func (s server) CreateTicket(ctx context.Context, request *kitchenpb.CreateTicketRequest) (*kitchenpb.CreateTicketResponse, error) {
	id := uuid.New().String()
	
	err := s.app.CreateTicket(ctx, commands.CreateTicket{
		ID:           id,
		RestaurantID: request.GetRestaurantID(),
		LineItems:    s.toLineItemsDomain(request.GetLineItems()),
	})
	if err != nil {
		return nil, err
	}
	return &kitchenpb.CreateTicketResponse{TicketID: id}, nil
}

func (s server) GetTicket(ctx context.Context, request *kitchenpb.GetTicketRequest) (*kitchenpb.GetTicketResponse, error) {
	ticket, err := s.app.GetTicket(ctx, queries.GetTicket{ID: request.GetTicketID()})
	if err != nil {
		return nil, err
	}
	return &kitchenpb.GetTicketResponse{
		RestaurantID: ticket.RestaurantID,
		LineItems:    s.toLineItemsProto(ticket.LineItems),
		Status:       ticket.Status.String(),
	}, nil
}

// func (s server) GetRestaurant(ctx context.Context, request *kitchenpb.GetRestaurantRequest) (*kitchenpb.GetRestaurantResponse, error) {
// }

// func (s server) AcceptTicket(ctx context.Context, request *kitchenpb.AcceptTicketRequest) (*kitchenpb.AcceptTicketResponse, error) {
// }

func (s server) toLineItemsDomain(lineItems []*kitchenpb.LineItem) []domain.LineItem {
	response := make([]domain.LineItem, 0, len(lineItems))
	for _, lineItem := range lineItems {
		response = append(response, domain.LineItem{
			MenuItemID: lineItem.MenuItemID,
			Name:       lineItem.Name,
			Quantity:   int(lineItem.Quantity),
		})
	}
	return response
}

func (s server) toLineItemsProto(lineItems []domain.LineItem) []*kitchenpb.LineItem {
	response := make([]*kitchenpb.LineItem, 0, len(lineItems))
	for _, lineItem := range lineItems {
		response = append(response, &kitchenpb.LineItem{
			MenuItemID: lineItem.MenuItemID,
			Name:       lineItem.Name,
			Quantity:   int64(lineItem.Quantity),
		})
	}
	return response
}
