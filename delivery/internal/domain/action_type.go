package domain

type ActionType string

const (
	UnknownAction ActionType = "Unknown"
	PickUp        ActionType = "PICKUP"
	DropOff       ActionType = "DROPOFF"
)

func (a ActionType) String() string {
	switch a {
	case PickUp, DropOff:
		return string(a)
	default:
		return ""
	}
}

func ToActionType(status string) ActionType {
	switch status {
	case PickUp.String():
		return PickUp
	case DropOff.String():
		return DropOff
	default:
		return UnknownAction
	}
}
