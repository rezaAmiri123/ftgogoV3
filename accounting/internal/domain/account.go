package domain

import "github.com/stackus/errors"

var (
	ErrAccountIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the account id cannot be blank")
	ErrAccountNameCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the account name cannot be blank")
	ErrAccountAlreadyEnabled    = errors.Wrap(errors.ErrBadRequest, "the account is already enabled")
	ErrAccountAlreadyDisabled   = errors.Wrap(errors.ErrBadRequest, "the account is already disabled")
)

type Account struct {
	ID      string
	Name    string
	Enabled bool
}

func RegisterAccount(id, name string) (*Account, error) {
	if id == "" {
		return nil, ErrAccountIDCannotBeBlank
	}
	if name == "" {
		return nil, ErrAccountNameCannotBeBlank
	}

	return &Account{
		ID:      id,
		Name:    name,
		Enabled: true,
	}, nil
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
