package strategy

import "fmt"

type UpiPayment struct {
	upiId string
}

func (u *UpiPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f amount using UPI ID %s\n", amount, u.upiId)
}