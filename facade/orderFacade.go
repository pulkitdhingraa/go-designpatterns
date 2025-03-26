package facade

type OrderFacade struct {
	inventory    *Inventory
	notification *Notification
}

func NewOrderFacade(inv *Inventory, notif *Notification) *OrderFacade {
	return &OrderFacade{
		inventory:    inv,
		notification: notif,
	}
}

func (o *OrderFacade) PlaceOrder(item string, qty int) {
	if _, err := o.inventory.ValidateInventory(item, qty); err != nil {
		panic(err)
	}
	o.notification.SendNotification()
}
