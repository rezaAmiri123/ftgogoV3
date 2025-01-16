package accountingpb

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
)

const (
	CommandChannel = "ftgogo.accounts.commands"

	AuthorizeAccountCommand = "accountingpb.AuthorizeAccount"
)

func Registration(reg registry.Registry) (err error) {
	serde := serdes.NewProtoSerde(reg)


	// Commands
	if err = serde.Register(&AuthorizeAccount{}); err != nil {
		return err
	}
	return nil
}

func (*AuthorizeAccount) Key() string { return AuthorizeAccountCommand }
