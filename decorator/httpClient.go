package decorator

import "fmt"

type HttpClient struct {
}

func (h *HttpClient) execute(clientId int) {
	fmt.Printf("Http enabled for client %d\n", clientId)
}