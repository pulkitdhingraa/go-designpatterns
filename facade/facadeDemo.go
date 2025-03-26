package facade

func Run() {
	user := NewUser("abc@xyz.com")
	inv := CreateInventory("book", 5)
	notif := NewNotification(user)
	order := NewOrderFacade(inv, notif)
	order.PlaceOrder("book", 2)
}