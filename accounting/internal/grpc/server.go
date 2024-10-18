package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/accountingpb"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/application"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	accountingpb.UnimplementedAccountingServiceServer
}

var _ accountingpb.AccountingServiceServer = (*server)(nil)

func RegisterServer(app application.App, register grpc.ServiceRegistrar) error {
	accountingpb.RegisterAccountingServiceServer(register, server{app: app})
	return nil
}

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
