package observer 

func Run() {
	reliance := NewRelianceStock(1700)
	retailer1 := &RetailObserver{
		userId: 1,
	}
	whaler1 := &WhalerObserver{
		userId: 2,
	}
	reliance.subscribe(retailer1)
	reliance.notifyAll()
	reliance.updatePrice(1750)
	reliance.subscribe(whaler1)
	reliance.notifyAll()
	reliance.unsubscribe(retailer1)
	reliance.updatePrice(1725)
	reliance.notifyAll()
}