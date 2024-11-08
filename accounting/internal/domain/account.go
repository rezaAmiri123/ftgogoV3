package domain

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/stackus/errors"
)

const AccountAggregate = "accounting.Account"

var (
	ErrAccountIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the account id cannot be blank")
	ErrAccountNameCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the account name cannot be blank")
	ErrAccountAlreadyEnabled    = errors.Wrap(errors.ErrBadRequest, "the account is already enabled")
	ErrAccountAlreadyDisabled   = errors.Wrap(errors.ErrBadRequest, "the account is already disabled")
	ErrAccountDisabled          = errors.Wrap(errors.ErrFailedPrecondition, "account is disabled")
)

type Account struct {
	ddd.Aggregate
	Name    string
	Enabled bool
}

func(Account)Key()string{return AccountAggregate}

func NewAccount(id string)*Account{
	return &Account{
		Aggregate: ddd.NewAggregate(id,AccountAggregate),
	}
}

func RegisterAccount(id, name string) (*Account, error) {
	if id == "" {
		return nil, ErrAccountIDCannotBeBlank
	}
	if name == "" {
		return nil, ErrAccountNameCannotBeBlank
	}

	account := NewAccount(id)
	account.Name = name
	account.Enabled = true

	return account, nil
}

func (a *Account) Enable() error {
	if a.Enabled {
		return ErrAccountAlreadyEnabled
	}
	a.Enabled = true
	return nil
}

func (a *Account) Disable() error {
	if !a.Enabled {
		return ErrAccountAlreadyDisabled
	}
	a.Enabled = false
	return nil
}

func (a *Account) AuthorizeOrder(OrderID string, OrderTotal int) error {
	if !a.Enabled {
		return ErrAccountDisabled
	}
	return nil
}
