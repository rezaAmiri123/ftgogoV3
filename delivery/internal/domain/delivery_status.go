package domain

type DeliveryStatus string

const (
	Unknown           DeliveryStatus = "Unknown"
	DeliveryPending   DeliveryStatus = "CreatePending"
	DeliveryScheduled DeliveryStatus = "AwaitingAcceptance"
	DeliveryCancelled DeliveryStatus = "Accepted"
)

func (s DeliveryStatus) String() string {
	switch s {
	case DeliveryPending, DeliveryScheduled, DeliveryCancelled:
		return string(s)
	default:
		return ""
	}
}

func ToDeliveryStatus(status string) DeliveryStatus {
	switch status {
	case DeliveryPending.String():
		return DeliveryPending
	case DeliveryScheduled.String():
		return DeliveryScheduled
	case DeliveryCancelled.String():
		return DeliveryCancelled
	default:
		return Unknown
	}
}
