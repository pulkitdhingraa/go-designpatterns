package strategy

type PaymentProcessor struct {
	strategy PaymentStrategy
}

// Strategy Setter
func (p *PaymentProcessor) SetPaymentStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentProcessor) ExecutePayment(amount float64) {
	p.strategy.Pay(amount)
}