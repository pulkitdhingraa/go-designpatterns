package observer 

type Stock interface {
	subscribe(user User)
	unsubscribe(user User)
	notifyAll()
}