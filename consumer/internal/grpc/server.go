package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/application"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	consumerpb.UnimplementedConsumerServiceServer
}

var _ consumerpb.ConsumerServiceServer = (*server)(nil)

func RegisterServer(app application.App, register grpc.ServiceRegistrar) error {
	consumerpb.RegisterConsumerServiceServer(register, server{app: app})
	return nil
}
func (s server) RegisterConsumer(ctx context.Context, request *consumerpb.RegisterConsumerRequest) (*consumerpb.RegisterConsumerResponse, error) {
	id := uuid.New()
	err := s.app.RegisterConsumer(ctx, application.RegisterConsumer{
		// ID: ,
	})
}
func (s server) GetConsumer(ctx context.Context, request *consumerpb.GetConsumerRequest) (*consumerpb.GetConsumerResponse, error)
func (s server) UpdateConsumer(ctx context.Context, request *consumerpb.UpdateConsumerRequest) (*consumerpb.UpdateConsumerResponse, error)
func (s server) GetAddress(ctx context.Context, request *consumerpb.GetAddressRequest) (*consumerpb.GetAddressResponse, error)
func (s server) UpdateAddress(ctx context.Context, request *consumerpb.UpdateAddressRequest) (*consumerpb.UpdateAddressResponse, error)
func (s server) RemoveAddress(ctx context.Context, request *consumerpb.RemoveAddressRequest) (*consumerpb.RemoveAddressResponse, error)

func (s server) GetAccount(ctx context.Context, request *accountingpb.GetAccountRequest) (*accountingpb.GetAccountResponse, error) {
	account, err := s.app.GetAccount(ctx, application.GetAccount{
		ID: request.GetAccountID(),
	})
	if err != nil {
		return nil, err
	}

	return &accountingpb.GetAccountResponse{
		AccountID: account.ID,
		Enabled:   account.Enabled,
	}, nil
}
func (s server) DisableAccount(ctx context.Context, request *accountingpb.DisableAccountRequest) (*accountingpb.DisableAccountResponse, error) {
	err := s.app.DisableAccount(ctx, application.DisableAccount{
		ID: request.GetAccountID(),
	})
	return &accountingpb.DisableAccountResponse{}, err
}

func (s server) EnableAccount(ctx context.Context, request *accountingpb.EnableAccountRequest) (*accountingpb.EnableAccountResponse, error) {
	err := s.app.EnableAccount(ctx, application.EnableAccount{
		ID: request.GetAccountID(),
	})
	return &accountingpb.EnableAccountResponse{}, err
}
