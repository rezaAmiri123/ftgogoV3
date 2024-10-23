package domain


import "context"

type CreateAccount struct{
	ID string
	Name string
}

type AccountRepository interface{
	CreateAccount(ctx context.Context, account CreateAccount)error
}