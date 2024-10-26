package domain

type Order struct {
	OrderID        string
	ConsumerID     string
	RestaurantID   string
	RestaurantName string
	Total          int
	Status         string
	// TODO EstimatedDelivery time.Duration
	// TODO other data
}
