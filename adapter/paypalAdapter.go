package adapter

type PayPalAdapter struct {
	PayPal *PayPalService
}

func NewPayPalAdapter(service *PayPalService) *PayPalAdapter {
	return &PayPalAdapter{
		PayPal: service,
	}
}

func (p *PayPalAdapter) ProcessPayment(amount float64, currency string){
	p.PayPal.Pay(amount, currency)
}