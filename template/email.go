package template 

import (
	"fmt"
)

type EmailNotification struct {
	Notification	// embedding
}

func (s *EmailNotification) Authenticate() {
	fmt.Println("Authenticating Email gateway...")
}

func (s *EmailNotification) FormatMessage(message string) string {
	return "Email: " + message
}

func (s *EmailNotification) Deliver(recipient string, message string) {
	fmt.Printf("Sending Email to %s: %s\n", recipient, message)
}