package domain

type LineItem struct {
	MenuItemID string
	Name       string
	Price      int
	Quantity   int
}
func (i LineItem) GetTotal() int {
	return i.Price * i.Quantity
}

type Address struct {
	Street1 string
	Street2 string
	City    string
	State   string
	Zip     string
}

type MenuItem struct {
	ID    string
	Name  string
	Price int
}

type MenuItemQuantities map[string]int