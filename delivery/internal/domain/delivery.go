package domain

import (
	"time"

	"github.com/stackus/errors"
)

var (
	ErrDeliveryIDCannotBeBlank      = errors.Wrap(errors.ErrBadRequest, "the delivery id cannot be blank")
	ErrRestaurantIDCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the restaurant id cannot be blank")
	ErrPickUpAddressCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the pick up address cannot be blank")
	ErrDeliveryAddressCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the delivery address cannot be blank")
)

type Delivery struct {
	ID                string
	RestaurantID      string
	AssignedCourierID string
	PickUpAddress     Address
	DeliveryAddress   Address
	Status            DeliveryStatus
	PickUpTime        time.Time
	ReadyBy           time.Time
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

	delivery := &Delivery{
		ID:                id,
		RestaurantID:      restaurantID,
		AssignedCourierID: "",
		PickUpAddress:     pickUpAddress,
		DeliveryAddress:   deliveryAddress,
		Status:            DeliveryPending,
	}
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
