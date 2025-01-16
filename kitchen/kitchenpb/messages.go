package kitchenpb

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"github.com/rezaAmiri123/ftgogoV3/internal/registry/serdes"
)

const (
	TicketAggregateChannel = "ftgogo.kitchens.events.Ticket"

	TicketAcceptedEvent = "kitchenpb.TicketAccepted"

	CommandChannel = "ftgogo.kitchens.commands"

	CreateTicketCommands        = "kitchenpb.CreateTicketCommands"
	ConfirmCreateTicketCommands = "kitchenpb.ConfirmCreateTicketCommands"
	CancelCreateTicketCommands  = "kitchenpb.CancelCreateTicketCommands"

	CreatedTicketReply = "kitchenpb.CreatedTicketReply"
)

func Registeration(reg registry.Registry) (err error) {
	serde := serdes.NewProtoSerde(reg)

	// Events
	if err = serde.Register(&TicketAccepted{}); err != nil {
		return err
	}

	// Commands
	if err = serde.Register(&CreateTicket{}); err != nil {
		return err
	}
	if err = serde.Register(&ConfirmCreateTicket{}); err != nil {
		return err
	}
	if err = serde.Register(&CancelCreateTicket{}); err != nil {
		return err
	}

	// Replies
	if err = serde.Register(&CreatedTicket{}); err != nil {
		return err
	}

	return nil
}

// Events
func (*TicketAccepted) Key() string { return TicketAcceptedEvent }

// commands
func (*CreateTicket) Key() string        { return CreateTicketCommands }
func (*ConfirmCreateTicket) Key() string { return ConfirmCreateTicketCommands }
func (*CancelCreateTicket) Key() string  { return CancelCreateTicketCommands }

// replies
func (*CreatedTicket) Key() string { return CreatedTicketReply }
