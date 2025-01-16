package commands

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
)

type ApproveOrder struct {
	ID       string
	TicketID string
}

type ApproveOrderHandler struct {
	orders      domain.OrderRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewApproveOrderHandler(
	orders domain.OrderRepository,
	publisher ddd.EventPublisher[ddd.Event],
) ApproveOrderHandler {
	return ApproveOrderHandler{
		orders:      orders,
		publisher: publisher,
	}
}

func (h ApproveOrderHandler) ApproveOrder(ctx context.Context, cmd ApproveOrder) error {
	order, err := h.orders.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}
	
	event, err := order.ApproveOrder(cmd.TicketID)
	if err != nil {
		return err
	}

	err = h.orders.Save(ctx, order)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
