package domain


import "context"

type CreateAccount struct{
	ID string
	Name string
}
//go:generate mockery --name AccountRepository
type AccountRepository interface{
	CreateAccount(ctx context.Context, account CreateAccount)error
}