package domain

import "github.com/stackus/errors"

// Restaurant errors
var (
	ErrRestaurantNotFound = errors.Wrap(errors.ErrNotFound, "restaurant not found")
)

type Restaurant struct {
	RestaurantID string
	Name         string
	Address      Address
	// note: no menu items
}
