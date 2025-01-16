package models

type CreateOrderData struct{
	OrderID string
	ConsumerID string
	RestaurantID string
	TicketID string
	LineItems []LineItem
	OrderTotal int
}

type LineItem struct {
	MenuItemID string
	Name       string
	Price      int
	Quantity   int
}

func (i LineItem) GetTotal() int {
	return i.Price * i.Quantity
}
