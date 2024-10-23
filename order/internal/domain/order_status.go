package domain

type OrderStatus string

const (
	UnknownOrderStatus OrderStatus = "UnknownOrderStatus"
	ApprovalPending    OrderStatus = "ApprovalPending"
	Approved           OrderStatus = "Approved"
	Rejected           OrderStatus = "Rejected"
	CancelPending      OrderStatus = "CancelPending"
	Cancelled          OrderStatus = "Cancelled"
	RevisionPending    OrderStatus = "RevisionPending"
)

func (s OrderStatus) String() string {
	switch s {
	case ApprovalPending, Approved, Rejected,
		CancelPending, Cancelled, RevisionPending:
		return string(s)
	default:
		return ""
	}
}

func ToOrderStatus(status string) OrderStatus {
	switch status {
	case ApprovalPending.String():
		return ApprovalPending
	case Approved.String():
		return Approved
	case Rejected.String():
		return Rejected
	case CancelPending.String():
		return CancelPending
	case Cancelled.String():
		return Cancelled
	case RevisionPending.String():
		return RevisionPending
	default:
		return UnknownOrderStatus
	}

}
