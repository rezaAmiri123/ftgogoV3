package logging

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger zerolog.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}

func (a Application) RegisterAccount(ctx context.Context, register application.RegisterAccount) (err error) {
	a.logger.Info().Msg("-->accounting.RegisterAccount")
	defer func() { a.logger.Info().Err(err).Msg("<--accounting.RegisterAccount") }()
	return a.App.RegisterAccount(ctx, register)
}

func (a Application) GetAccount(ctx context.Context, get application.GetAccount) (_ *domain.Account, err error) {
	a.logger.Info().Msg("-->accounting.GetAccount")
	defer func() { a.logger.Info().Err(err).Msg("<--accounting.GetAccount") }()
	return a.App.GetAccount(ctx, get)
}
func (a Application) EnableAccount(ctx context.Context, enable application.EnableAccount) (err error) {
	a.logger.Info().Msg("-->accounting.EnableAccount")
	defer func() { a.logger.Info().Err(err).Msg("<--accounting.EnableAccount") }()
	return a.App.EnableAccount(ctx, enable)

}
func (a Application) DisableAccount(ctx context.Context, disable application.DisableAccount) (err error) {
	a.logger.Info().Msg("-->accounting.DisableAccount")
	defer func() { a.logger.Info().Err(err).Msg("<--accounting.DisableAccount") }()
	return a.App.DisableAccount(ctx, disable)
}
