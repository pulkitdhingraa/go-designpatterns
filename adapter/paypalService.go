package adapter

import "fmt"

type PayPalService struct{}

func (p *PayPalService) Pay(amount float64, currency string) {
	fmt.Printf("Paid %.2f %s of amount with PayPal", amount, currency)
}