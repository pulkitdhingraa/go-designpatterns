package decorator

import "fmt"

type Logging struct {
	client HttpClientInterface
}

func (l *Logging) execute(clientId int) {
	fmt.Printf("Logging enabled for client %d\n", clientId)
	l.client.execute(clientId)
}