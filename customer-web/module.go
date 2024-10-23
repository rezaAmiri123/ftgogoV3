package customerweb

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/rest"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	var app application.App
	// app = application.New(consumers, accounts)
	// app = logging.LogApplicationAccess(app, mono.Logger())

	server := rest.NewServer(app)
	if err := rest.RegisterServer(server,mono.Mux()); err != nil {
		return err
	}

	return nil
}
