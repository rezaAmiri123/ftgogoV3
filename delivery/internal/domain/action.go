package domain

import "time"

type Action struct {
	DeliveryID string
	ActionType ActionType
	Address    Address
	When       time.Time
}

func (a Action) IsFor(deliveryID string) bool {
	return a.DeliveryID == deliveryID
}

type Plan []Action

func (p *Plan) Add(action Action) {
	*p = append(*p, action)
}

func (p *Plan) RemoveDelivery(deliveryID string) {
	replacement := Plan{}
	for _, action := range *p {
		if !action.IsFor(deliveryID) {
			replacement = append(replacement, action)
		}
	}

	*p = replacement
}

func (p Plan) ActionFor(deliveryID string) []Action {
	actions := []Action{}
	for _, action := range p {
		if action.IsFor(deliveryID) {
			actions = append(actions, action)
		}
	}
	return actions
}
