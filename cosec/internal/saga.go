package internal

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/accountingpb"
	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/cosec/internal/models"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/sec"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/kitchenpb"
	"github.com/rezaAmiri123/ftgogoV3/order/orderpb"
)

const (
	CreateOrderSagaName     = "cosec.CreateOrder"
	CreateOrderReplyChannel = "ftgogo.cosec.replies.CreateOrder"
)

type createOrderSaga struct {
	sec.Saga[*models.CreateOrderData]
}

func NewCreateOrderSaga() sec.Saga[*models.CreateOrderData] {
	saga := createOrderSaga{
		Saga: sec.NewSaga[*models.CreateOrderData](CreateOrderSagaName, CreateOrderReplyChannel),
	}

	// 0. RejectOrder
	saga.AddStep().Compensation(saga.rejectOrder) // TODO this method has not been implemented on order yet

	// 1. AuthorizeConsumer
	saga.AddStep().
		Action(saga.authorizeConsumer)

	// 2. Create ticket
	saga.AddStep().
		Action(saga.createTicket).
		OnActionReply(kitchenpb.CreatedTicketReply, saga.onCreatedTicketReply).
		Compensation(saga.cancelCreateTicket) // TODO this method has not been implemented on kitchen yet
	// 3. AuthorizeAccount
	saga.AddStep().
		Action(saga.authorizeAccount)

	// 4. Confirm Create ticket
	saga.AddStep().
		Action(saga.confirmCreateTicket)

	// 5. Approve Order
	saga.AddStep().
		Action(saga.approveOrder)

	return saga
}

func (s createOrderSaga) rejectOrder(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	cmd := ddd.NewCommand(orderpb.RejectOrderCommand, &orderpb.RejectOrder{
		OrderID: data.OrderID,
	})
	return orderpb.CommandChannel, cmd, nil
}

func (s createOrderSaga) authorizeConsumer(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	cmd := ddd.NewCommand(consumerpb.AuthorizeConsumerCommand, &consumerpb.AuthorizeCustomer{
		Id:         data.ConsumerID,
		OrderId:    data.OrderID,
		TotalOrder: int64(data.OrderTotal),
	})

	return consumerpb.CommandChannel, cmd, nil
}

func (s createOrderSaga) createTicket(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	items := make([]*kitchenpb.CreateTicket_LineItem, len(data.LineItems))
	for i, item := range data.LineItems {
		items[i] = &kitchenpb.CreateTicket_LineItem{
			MenuItemID: item.MenuItemID,
			Name:       item.Name,
			Quantity:   int64(item.Quantity),
		}
	}

	cmd := ddd.NewCommand(kitchenpb.CreateTicketCommands, &kitchenpb.CreateTicket{
		OrderID:      data.OrderID,
		RestaurantID: data.RestaurantID,
		Items:        items,
	})

	return kitchenpb.CommandChannel, cmd, nil
}

func (s createOrderSaga) onCreatedTicketReply(ctx context.Context, data *models.CreateOrderData, reply ddd.Reply) error {
	payload := reply.Payload().(*kitchenpb.CreatedTicket)

	data.TicketID = payload.GetId()

	return nil
}

func (s createOrderSaga) cancelCreateTicket(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	cmd := ddd.NewCommand(kitchenpb.CancelCreateTicketCommands, &kitchenpb.CancelCreateTicket{
		TicketID: data.TicketID,
	})

	return kitchenpb.CommandChannel, cmd, nil
}

func (s createOrderSaga) authorizeAccount(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	cmd := ddd.NewCommand(accountingpb.AuthorizeAccountCommand, &accountingpb.AuthorizeAccount{
		AccountId:  data.ConsumerID,
		OrderId:    data.OrderID,
		TotalOrder: int64(data.OrderTotal),
	})

	return accountingpb.CommandChannel, cmd, nil
}

func (s createOrderSaga) confirmCreateTicket(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	cmd := ddd.NewCommand(kitchenpb.ConfirmCreateTicketCommands, &kitchenpb.ConfirmCreateTicket{
		TicketID: data.TicketID,
	})

	return kitchenpb.CommandChannel, cmd, nil
}

func (s createOrderSaga) approveOrder(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	cmd := ddd.NewCommand(orderpb.ApproveOrderCommand, &orderpb.ApproveOrder{
		OrderID:  data.OrderID,
		TicketID: data.TicketID,
	})

	return orderpb.CommandChannel, cmd, nil
}
