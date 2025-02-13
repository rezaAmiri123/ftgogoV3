package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/rezaAmiri123/ftgogoV3/internal/errorsotel"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/application/queries"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	orderpb.UnimplementedOrderServiceServer
}

var _ orderpb.OrderServiceServer = (*server)(nil)

func RegisterServer(app application.App, register grpc.ServiceRegistrar) error {
	orderpb.RegisterOrderServiceServer(register, server{app: app})
	return nil
}

func (s server) CreateOrder(ctx context.Context, request *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	span := trace.SpanFromContext(ctx)
	id := uuid.New().String()

	span.SetAttributes(
		attribute.String("OrderID", id),
		attribute.String("ConsumerID", request.GetConsumerID()),
		attribute.String("RestaurantID", request.GetRestaurantID()),
	)

	lineItems := make(map[string]int, len(request.LineItems.Items))
	for k, v := range request.LineItems.Items {
		lineItems[k] = int(v)
	}

	err := s.app.CreateOrder(ctx, commands.CreateOrder{
		ID:           id,
		RestaurantID: request.GetRestaurantID(),
		ConsumerID:   request.GetConsumerID(),
		DeliverTo:    s.toAddressDomain(request.GetDeliverTo()),
		DeliverAt:    request.DeliverAt.AsTime(),
		LineItems:    lineItems,
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return &orderpb.CreateOrderResponse{OrderID: id}, nil
}

func (s server) GetOrder(ctx context.Context, request *orderpb.GetOrderRequest) (*orderpb.GetOrderResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetOrderID()),
	)

	order, err := s.app.GetOrder(ctx, queries.GetOrder{ID: request.GetOrderID()})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	
	return &orderpb.GetOrderResponse{
		Order: s.toOrderProto(order),
	}, nil
}

func (s server) toOrderProto(order *domain.Order) *orderpb.Order {
	return &orderpb.Order{
		OrderID:      order.ID(),
		ConsumerID:   order.ConsumerID,
		RestaurantID: order.RestaurantID,
		Status:       order.Status.String(),
		OrderTotal:   int64(order.OrderTotal()),
	}
}

func (s server) toAddressDomain(address *orderpb.Address) domain.Address {
	return domain.Address{
		Street1: address.Street1,
		Street2: address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}
