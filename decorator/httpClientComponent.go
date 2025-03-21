package decorator

type HttpClientInterface interface {
	execute(clientId int)
}