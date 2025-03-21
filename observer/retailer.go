package observer

import "fmt"

type RetailObserver struct {
	userId int
}

func (r *RetailObserver) update(price float64) {
	fmt.Printf("Updating retail user %d about new price %.2f\n", r.userId, price)
}

func (r *RetailObserver) getUserId() int {
	return r.userId
}