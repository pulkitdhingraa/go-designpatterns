package observer 

import "fmt"

type WhalerObserver struct {
	userId int
}

func (w *WhalerObserver) update(price float64) {
	fmt.Printf("Updating whaler user %d about new price %.2f\n", w.userId, price)
}

func (w *WhalerObserver) getUserId() int {
	return w.userId
}