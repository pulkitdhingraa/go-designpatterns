package singleton

import (
	"fmt"
	"sync"
)

type logger struct {
	log string
}

var loggerInstance *logger
var mu = &sync.Mutex{}

func getLoggerInstance(wg *sync.WaitGroup) *logger{
	defer wg.Done()
	// avoid unncessary locking
	if loggerInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		// check again whether another goroutine created instance between the time of first check and acquiring lock
		if loggerInstance == nil {
			loggerInstance = &logger{}
			fmt.Println("New logger instance created")
		} else {
			fmt.Println("Logger instance already exists")
		}
	} else {
		fmt.Println("Logger instance already exists")
	}
	return loggerInstance
}