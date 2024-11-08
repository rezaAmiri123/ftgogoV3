package commands

import (
	"context"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
)

type CreateOrder struct {
	ID           string
	ConsumerID   string
	RestaurantID string
	DeliverAt    time.Time
	DeliverTo    domain.Address
	LineItems    map[string]int
}

type CreateOrderHandler struct {
	orders      domain.OrderRepository
	restaurants domain.RestaurantRepository
	consumers   domain.ConsumerRepository
	kitchens    domain.KitchenRepository
	accounts    domain.AccountRepository
}

func NewCreateOrderHandler(
	orders domain.OrderRepository,
	restaurants domain.RestaurantRepository,
	consumers domain.ConsumerRepository,
	kitchens domain.KitchenRepository,
	accounts domain.AccountRepository,
) CreateOrderHandler {
	return CreateOrderHandler{
		orders:      orders,
		restaurants: restaurants,
		consumers:   consumers,
		kitchens:    kitchens,
		accounts:    accounts,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	order, err := h.orders.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}
	restaurant, err := h.restaurants.Find(ctx, cmd.RestaurantID)
	if err != nil {
		return err
	}
	lineItems, err := restaurant.LineItems(cmd.LineItems)
	if err != nil {
		return err
	}
	err = order.CreateOrder(cmd.ConsumerID, cmd.RestaurantID, lineItems, cmd.DeliverAt, cmd.DeliverTo)
	if err != nil {
		return err
	}

	// Validate order by consumer
	err = h.consumers.ValidateOrderByConsumer(ctx, domain.ValidateOrderByConsumer{
		ConsumerID: cmd.ConsumerID,
		OrderID:    cmd.ID,
		OrderTotal: order.OrderTotal(),
	})
	if err != nil {
		return err
	}

	// Create Ticket
	err = h.kitchens.CreateTicket(ctx, domain.CreateTicket{
		ID:           cmd.ID,
		RestaurantID: cmd.RestaurantID,
		TicketDetail: lineItems,
	})
	if err != nil {
		return err
	}

	// Authorize order by account
	err = h.accounts.AuthorizeOrderByAccount(ctx, domain.AuthorizeOrderByAccount{
		AccountID:  cmd.ConsumerID,
		OrderID:    cmd.ID,
		OrderTotal: order.OrderTotal(),
	})
	if err != nil {
		return err
	}

	// Confirm Create ticket
	err = h.kitchens.ConfirmCreateTicket(ctx, order.ID())
	if err != nil {
		return err
	}

	// Approve Order
	err = order.ApproveOrder(order.ID())
	if err != nil {
		return err
	}

	err = h.orders.Save(ctx, order)
	if err != nil {
		return err
	}
	return nil
}
