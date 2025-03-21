package observer

type RelianceStock struct {
	observerList []User
	price float64
}

func NewRelianceStock(price float64) *RelianceStock{
	return &RelianceStock{
		observerList: make([]User, 0),
		price: price,
	}
}

func (r *RelianceStock) subscribe(user User) {
	r.observerList = append(r.observerList, user)
}

func (r *RelianceStock) unsubscribe(user User) {
	r.observerList = removeFromObserverList(r.observerList, user)
}

func (r *RelianceStock) notifyAll() {
	for _, observer := range r.observerList {
		observer.update(r.price)
	}
}

func (r *RelianceStock) updatePrice(newPrice float64) {
	r.price = newPrice
}

func removeFromObserverList(observerList []User, user User) []User {
	for i,v := range observerList {
		if v.getUserId() == user.getUserId() {
			return append(observerList[:i], observerList[i+1:]...)
		}
	}
	return observerList
}