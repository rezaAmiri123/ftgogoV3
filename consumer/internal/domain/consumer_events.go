package domain

type ConsumerRegistered struct{
	Consumer *Consumer
}

func (ConsumerRegistered)EventName()string{return "consumers.ConsumerRegistered"}
