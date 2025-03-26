package facade

import "fmt"

type Notification struct {
	user *User
}

func NewNotification(user *User) *Notification {
	return &Notification{
		user: user,
	}
}

func (n *Notification) SendNotification() {
	fmt.Printf("Notification sent to user %s", n.user.email)
}