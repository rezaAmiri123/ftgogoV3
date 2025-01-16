package internal

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/accountingpb"
	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/cosec/internal/models"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
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

func (s createOrderSaga) rejectOrder(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(orderpb.RejectOrderCommand, orderpb.CommandChannel,
		&orderpb.RejectOrder{
			OrderID: data.OrderID,
		})
}

func (s createOrderSaga) authorizeConsumer(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(consumerpb.AuthorizeConsumerCommand, consumerpb.CommandChannel,
		&consumerpb.AuthorizeCustomer{
			Id:         data.ConsumerID,
			OrderId:    data.OrderID,
			TotalOrder: int64(data.OrderTotal),
		})
}

func (s createOrderSaga) createTicket(ctx context.Context, data *models.CreateOrderData) am.Command {
	items := make([]*kitchenpb.CreateTicket_LineItem, len(data.LineItems))
	for i, item := range data.LineItems {
		items[i] = &kitchenpb.CreateTicket_LineItem{
			MenuItemID: item.MenuItemID,
			Name:       item.Name,
			Quantity:   int64(item.Quantity),
		}
	}

	return am.NewCommand(kitchenpb.CreateTicketCommands, kitchenpb.CommandChannel, &kitchenpb.CreateTicket{
		OrderID:      data.OrderID,
		RestaurantID: data.RestaurantID,
		Items:        items,
	})
}

func (s createOrderSaga) onCreatedTicketReply(ctx context.Context, data *models.CreateOrderData, reply ddd.Reply) error {
	payload := reply.Payload().(*kitchenpb.CreatedTicket)

	data.TicketID = payload.GetId()

	return nil
}

func (s createOrderSaga) cancelCreateTicket(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(kitchenpb.CancelCreateTicketCommands, kitchenpb.CommandChannel,
		&kitchenpb.CancelCreateTicket{
			TicketID: data.TicketID,
		})
}

func (s createOrderSaga) authorizeAccount(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(accountingpb.AuthorizeAccountCommand, accountingpb.CommandChannel,
		&accountingpb.AuthorizeAccount{
			AccountId:  data.ConsumerID,
			OrderId:    data.OrderID,
			TotalOrder: int64(data.OrderTotal),
		})
}

func (s createOrderSaga) confirmCreateTicket(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(kitchenpb.ConfirmCreateTicketCommands, kitchenpb.CommandChannel,
		&kitchenpb.ConfirmCreateTicket{
			TicketID: data.TicketID,
		})
}

func (s createOrderSaga) approveOrder(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(orderpb.ApproveOrderCommand, orderpb.CommandChannel,
		&orderpb.ApproveOrder{
			OrderID:  data.OrderID,
			TicketID: data.TicketID,
		})
}
