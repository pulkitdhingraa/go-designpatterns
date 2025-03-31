package template 

func Run() {
	// example using delegation
	sms := &SMSNotification{}
	n := Notification{
		inoitif: sms,
	}
	n.Send("Hi, How are you?", "abc@gmail.com")

	// example using embedding
	email := &EmailNotification{}
	Send2(email, "Lets meet tonight", "xyz@gmail.com")
}