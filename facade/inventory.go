package facade

import ("errors")

type Inventory struct {
	item string
	qty  int
}

func CreateInventory(item string, qty int) *Inventory {
	return &Inventory{
		item: "book",
		qty: 5,
	}
}

func (i *Inventory) ValidateInventory(item string, qty int) (bool, error) {
	if item == i.item && qty <= i.qty {
		return true, nil
	}
	return false, errors.New("not enough qty in inventory")
}