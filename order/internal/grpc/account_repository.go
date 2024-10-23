package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/accountingpb"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
	"google.golang.org/grpc"
)

type AccountRepository struct {
	client accountingpb.AccountingServiceClient
}

var _ domain.AccountRepository = (*AccountRepository)(nil)

func NewAccountRepository(conn *grpc.ClientConn) AccountRepository {
	return AccountRepository{client: accountingpb.NewAccountingServiceClient(conn)}
}

func (r AccountRepository) AuthorizeOrderByAccount(ctx context.Context, validate domain.AuthorizeOrderByAccount) error {
	_, err := r.client.AuthorizeOrderByAccount(ctx, &accountingpb.AuthorizeOrderByAccountRequest{
		AccountID:  validate.AccountID,
		OrderID:    validate.OrderID,
		OrderTotal: int64(validate.OrderTotal),
	})
	return err
}
