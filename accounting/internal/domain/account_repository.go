package domain

import "context"

type AccountRepository interface{
	Save(ctx context.Context, account *Account)error
	Find(ctx context.Context, accountID string)(*Account,error)
	Update(ctx context.Context,account *Account)error
}