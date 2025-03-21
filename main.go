package main

import (
	"go-designpatterns/singleton"
	"go-designpatterns/strategy"
	"go-designpatterns/decorator"
	"go-designpatterns/observer"
)

func main() {
	singleton.Run()
	strategy.Run()
	decorator.Run()
	observer.Run()
}