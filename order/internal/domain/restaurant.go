package domain

import "github.com/stackus/errors"

var (
	ErrMenuItemNotFound = errors.Wrap(errors.ErrNotFound, "menu item not found")
)

type Restaurant struct {
	ID        string
	Name      string
	MenuItems []MenuItem
}

func (r Restaurant) FindMenuItem(menuItemID string) (MenuItem, error) {
	for _, item := range r.MenuItems {
		if menuItemID == item.ID {
			return item, nil
		}
	}
	return MenuItem{}, errors.Wrap(ErrMenuItemNotFound, menuItemID)
}

func (r Restaurant) LineItems(mapLineItems map[string]int) ([]LineItem, error) {
	lineItems := make([]LineItem, 0, len(mapLineItems))
	for menuItemID, quantity := range mapLineItems {
		menuItem, err := r.FindMenuItem(menuItemID)
		if err != nil {
			return []LineItem{}, err
		}
		lineItems = append(lineItems, LineItem{
			MenuItemID: menuItemID,
			Name:       menuItem.Name,
			Price:      menuItem.Price,
			Quantity:   quantity,
		})
	}
	return lineItems, nil
}
