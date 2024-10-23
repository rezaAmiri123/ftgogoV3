package domain

import (
	"context"
)

type AuthorizeOrderByAccount struct {
	AccountID string
	OrderID    string
	OrderTotal int
}

type AccountRepository interface {
	AuthorizeOrderByAccount(ctx context.Context, validate AuthorizeOrderByAccount) error
}
