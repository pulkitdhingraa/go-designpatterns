package decorator

func Run() {
	client1 := &HttpClient{}
	// client1.execute(1)
	client1WithLogging := &Logging{client: client1}
	// client1WithLogging.execute(1)
	client1WithLoggingAndMetrics := &Metrics{client: client1WithLogging}
	client1WithLoggingAndMetrics.execute(1)
}