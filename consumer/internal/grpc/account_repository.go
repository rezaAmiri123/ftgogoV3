package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/accountingpb"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"google.golang.org/grpc"
)

type AccountRepository struct {
	client accountingpb.AccountingServiceClient
}

var _ domain.AccountRepository = (*AccountRepository)(nil)

func NewAccountRepository(conn *grpc.ClientConn) AccountRepository {
	return AccountRepository{client: accountingpb.NewAccountingServiceClient(conn)}
}

func (r AccountRepository) CreateAccount(ctx context.Context, account domain.CreateAccount) error {
	_, err := r.client.CreateAccount(ctx, &accountingpb.CreateAccountRequest{
		ID:   account.ID,
		Name: account.Name,
	})
	return err
}
