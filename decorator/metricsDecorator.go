package decorator

import "fmt"

type Metrics struct {
	client HttpClientInterface
}

func (m *Metrics) execute(clientId int) {
	fmt.Printf("Metrics enabled for client %d\n", clientId)
	m.client.execute(clientId)
}