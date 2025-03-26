package adapter

func Run() {
	payPalService := &PayPalService{}
	payPalAdapter := NewPayPalAdapter(payPalService)
	payPalAdapter.ProcessPayment(12.50, "USD")
}