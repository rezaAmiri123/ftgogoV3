package domain

import (
	"time"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/stackus/errors"
)

const DeliveryAggregate = "delivery.DeliveryAggregate"

var (
	ErrDeliveryIDCannotBeBlank      = errors.Wrap(errors.ErrBadRequest, "the delivery id cannot be blank")
	ErrRestaurantIDCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the restaurant id cannot be blank")
	ErrPickUpAddressCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the pick up address cannot be blank")
	ErrDeliveryAddressCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the delivery address cannot be blank")
)

type Delivery struct {
	ddd.Aggregate
	RestaurantID      string
	AssignedCourierID string
	PickUpAddress     Address
	DeliveryAddress   Address
	Status            DeliveryStatus
	PickUpTime        time.Time
	ReadyBy           time.Time
}

func (Delivery) key() string { return DeliveryAggregate }

func NewDelivery(id string) *Delivery {
	return &Delivery{
		Aggregate: ddd.NewAggregate(id, DeliveryAggregate),
	}
}

func CreateDelivery(id, restaurantID string, pickUpAddress, deliveryAddress Address) (*Delivery, error) {
	if id == "" {
		return nil, ErrDeliveryIDCannotBeBlank
	}
	if restaurantID == "" {
		return nil, ErrRestaurantIDCannotBeBlank
	}
	if pickUpAddress == (Address{}) {
		return nil, ErrPickUpAddressCannotBeBlank
	}
	if deliveryAddress == (Address{}) {
		return nil, ErrDeliveryAddressCannotBeBlank
	}

	delivery :=NewDelivery(id)
	delivery.RestaurantID = restaurantID
	delivery.AssignedCourierID = ""
	delivery.PickUpAddress = pickUpAddress
	delivery.DeliveryAddress = deliveryAddress
	delivery.Status = DeliveryPending

	return delivery, nil
}

func (d *Delivery) Schedule(readyBy time.Time, assignedCourierID string) {
	d.ReadyBy = readyBy
	d.AssignedCourierID = assignedCourierID
	d.Status = DeliveryScheduled
}

func (d *Delivery) Cancel() {
	d.AssignedCourierID = ""
	d.Status = DeliveryCancelled
}
