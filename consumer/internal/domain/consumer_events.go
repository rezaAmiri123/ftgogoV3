package domain

const (
	ConsumerRegisteredEvent = "consumers.ConsumerRegistered"
)

type ConsumerRegistered struct {
	Consumer *Consumer
}

func (ConsumerRegistered) Key() string { return ConsumerRegisteredEvent }
