package domain

import (
	"context"
	"time"
)

type (
	AcceptTicket struct {
		ID      string
		ReadyBy time.Time
	}
)

type KitchenRepository interface {
	AcceptTicket(ctx context.Context, accept AcceptTicket) error
}
