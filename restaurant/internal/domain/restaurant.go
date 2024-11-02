package domain

import (
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/stackus/errors"
)

var (
	ErrRestaurantIDCannotBeBlank      = errors.Wrap(errors.ErrBadRequest, "the restaurant id cannot be blank")
	ErrRestaurantNameCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the restaurant name cannot be blank")
	ErrRestaurantAddressCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the restaurant address cannot be blank")
	ErrMenuItemNotFound               = errors.Wrap(errors.ErrNotFound, "the menu item not found")
)

type Restaurant struct {
	ddd.AggregateBase
	Name      string
	Address   Address
	MenuItems map[string]MenuItem
}

func CreateRestaurant(id, name string, address Address) (*Restaurant, error) {
	if id == "" {
		return nil, ErrRestaurantIDCannotBeBlank
	}
	if name == "" {
		return nil, ErrRestaurantNameCannotBeBlank
	}
	if address == (Address{}) {
		return nil, ErrRestaurantAddressCannotBeBlank
	}
	restaurant := &Restaurant{
		AggregateBase: ddd.AggregateBase{ID: id},
		Name:          name,
		Address:       address,
		MenuItems:     make(map[string]MenuItem),
	}
	return restaurant, nil
}

func (r *Restaurant) FindMenuItem(menuItemID string) (MenuItem, error) {
	menuItem, ok := r.MenuItems[menuItemID]
	if !ok {
		return MenuItem{}, ErrMenuItemNotFound
	}
	return menuItem, nil
}

func (r *Restaurant) UpdateMenuItem(menuItems []MenuItem) error {
	for _, menuItem := range menuItems {
		r.MenuItems[menuItem.ID] = menuItem
	}
	return nil
}
