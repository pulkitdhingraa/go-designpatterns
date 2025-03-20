package strategy

func Run() {
	upiObj := UpiPayment{
		upiId: "abc@ybl",
	}
	// creditObj := CreditCard{
	// 	cardHolder: "abc",
	// 	cardNumber: "123-xxx",
	// 	cardExpiry: "04/28",
	// }
	processor := PaymentProcessor {
		strategy: &upiObj,
	}
	processor.ExecutePayment(12.25)
}