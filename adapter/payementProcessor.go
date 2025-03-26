package adapter

type PayementProcessor interface {
	ProcessPayment(amount float64, currency string)
}