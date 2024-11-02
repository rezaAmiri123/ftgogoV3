package domain

type TicketAccepted struct{
	Ticket *Ticket
} 

func(TicketAccepted)EventName()string{return "kitchen.TicketAccepted"}