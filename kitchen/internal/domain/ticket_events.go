package domain

const (
	TicketAcceptedEvent = "kitchen.TicketAccepted"
)

type TicketAccepted struct {
	Ticket *Ticket
}

func (TicketAccepted) Key() string { return "kitchen.TicketAccepted" }
