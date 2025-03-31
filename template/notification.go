package template

import "fmt"

type INotification interface {
	Authenticate()
	FormatMessage(message string) string
	Deliver(recipient string, message string)
}

type Notification struct {
	inoitif INotification	// delegation, define different notification types at runtime
}

func (n *Notification) Send(message string, recipient string) {
	n.inoitif.Authenticate()
	formattedMessage := n.inoitif.FormatMessage(message)
	n.inoitif.Deliver(recipient, formattedMessage)
	fmt.Println("Notification sent successfully")
}

func Send2(n INotification, message string, recipient string) {
	n.Authenticate()
	formattedMsg := n.FormatMessage(message)
	n.Deliver(recipient, formattedMsg)
}