package strategy

import (
	"fmt"
)

type CreditCard struct {
	cardHolder string
	cardNumber string
	cardExpiry string
}

func (c *CreditCard) Pay(amount float64) {
	fmt.Printf("Processing credit card payment of $%.2f for card holder %s\n", amount, c.cardHolder)
}