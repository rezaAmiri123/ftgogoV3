package application

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/domain"
)

type (
	RegisterAccount struct {
		ID   string
		Name string
	}
	GetAccount struct {
		ID string
	}
	EnableAccount struct {
		ID string
	}
	DisableAccount struct {
		ID string
	}
	AuthorizeOrderByAccount struct {
		ID         string
		OrderID    string
		OrderTotal int
	}
	App interface {
		RegisterAccount(ctx context.Context, register RegisterAccount) error
		GetAccount(ctx context.Context, get GetAccount) (*domain.Account, error)
		EnableAccount(ctx context.Context, enable EnableAccount) error
		DisableAccount(ctx context.Context, disable DisableAccount) error
		AuthorizeOrderByAccount(ctx context.Context, authorize AuthorizeOrderByAccount) error
	}

	Application struct {
		Account domain.AccountRepository
	}
)

var _ App = (*Application)(nil)

func New(account domain.AccountRepository) *Application {
	return &Application{
		Account: account,
	}
}

func (a Application) RegisterAccount(ctx context.Context, register RegisterAccount) error {
	account, err := domain.RegisterAccount(register.ID, register.Name)
	if err != nil {
		return err
	}

	return a.Account.Save(ctx, account)
}

func (a Application) GetAccount(ctx context.Context, get GetAccount) (*domain.Account, error) {
	return a.Account.Find(ctx, get.ID)
}

func (a Application) EnableAccount(ctx context.Context, enable EnableAccount) error {
	account, err := a.Account.Find(ctx, enable.ID)
	if err != nil {
		return err
	}

	err = account.Enable()
	if err != nil {
		return err
	}

	return a.Account.Update(ctx, account)
}

func (a Application) DisableAccount(ctx context.Context, disable DisableAccount) error {
	account, err := a.Account.Find(ctx, disable.ID)
	if err != nil {
		return err
	}

	err = account.Disable()
	if err != nil {
		return err
	}

	return a.Account.Update(ctx, account)
}

func (a Application) AuthorizeOrderByAccount(ctx context.Context, authorize AuthorizeOrderByAccount) error {
	account, err := a.Account.Find(ctx, authorize.ID)
	if err != nil {
		return err
	}

	return account.AuthorizeOrder(authorize.OrderID, authorize.OrderTotal)
}
