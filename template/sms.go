package template 

import (
	"fmt"
)

type SMSNotification struct {}

func (s *SMSNotification) Authenticate() {
	fmt.Println("\nAuthenticating SMS gateway...")
}

func (s *SMSNotification) FormatMessage(message string) string {
	return "SMS: " + message
}

func (s *SMSNotification) Deliver(recipient string, message string) {
	fmt.Printf("Sending SMS to %s: %s\n", recipient, message)
}