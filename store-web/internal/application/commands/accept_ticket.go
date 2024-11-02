package commands

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/domain"
)

type AcceptTicket struct {
	ID      string
	ReadyBy time.Time
}

type AcceptTicketHandler struct {
	kitchens domain.KitchenRepository
}

func NewAcceptTicketHandler(kitchens domain.KitchenRepository) AcceptTicketHandler {
	return AcceptTicketHandler{
		kitchens: kitchens,
	}
}

func (h AcceptTicketHandler) AcceptTicket(ctx context.Context, cmd AcceptTicket) error {
	err := h.kitchens.AcceptTicket(ctx, domain.AcceptTicket{
		ID:      cmd.ID,
		ReadyBy: cmd.ReadyBy,
	})
	return err
}
