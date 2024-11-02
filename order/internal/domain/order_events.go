package domain

type OrderCreated struct{
	Order *Order
}

func(OrderCreated)EventName()string{ return "order.OrderCreated"}