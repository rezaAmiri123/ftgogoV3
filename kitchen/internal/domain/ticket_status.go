package domain

type TicketStatus string

const (
	Unknown            TicketStatus = "Unknown"
	CreatePending      TicketStatus = "CreatePending"
	AwaitingAcceptance TicketStatus = "AwaitingAcceptance"
	Accepted           TicketStatus = "Accepted"
	Preparing          TicketStatus = "Preparing"
	ReadyForPickup     TicketStatus = "ReadyForPickup"
	PickedUp           TicketStatus = "PickedUp"
	CancelPending      TicketStatus = "CancelPending"
	Cancelled          TicketStatus = "Cancelled"
	RevisionPending    TicketStatus = "RevisionPending"
)

func (s TicketStatus) String() string {
	switch s {
	case CreatePending, AwaitingAcceptance, Accepted, Preparing, ReadyForPickup,
		PickedUp, CancelPending, Cancelled, RevisionPending:
		return string(s)
	default:
		return ""
	}
}

func ToTicketStatus(status string) TicketStatus {
	switch status {
	case CreatePending.String():
		return CreatePending
	case AwaitingAcceptance.String():
		return AwaitingAcceptance
	case Accepted.String():
		return Accepted
	case Preparing.String():
		return Preparing
	case ReadyForPickup.String():
		return ReadyForPickup
	case PickedUp.String():
		return PickedUp
	case CancelPending.String():
		return CancelPending
	case Cancelled.String():
		return Cancelled
	case RevisionPending.String():
		return RevisionPending
	default:
		return Unknown
	}
}
