package singleton

import (
	"sync"
)

func Run() {
	var wg sync.WaitGroup

	// multiple go routines trying to create logger instance
	for i:=0;i<5;i++ {
		wg.Add(1)
		go getLoggerInstance(&wg)
		// go getLoggerInstanceOnce(&wg)
	}

	wg.Wait()
}