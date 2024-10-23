package domain

type Courier struct {
	ID        string
	Plan      Plan
	Available bool
}

func(c *Courier)AddAction(action Action){
	c.Plan.Add(action)
}

func(c *Courier)CancelDelivery(deliveryID string){
	c.Plan.RemoveDelivery(deliveryID)
}
