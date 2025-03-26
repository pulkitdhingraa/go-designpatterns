package main

import (
	"go-designpatterns/singleton"
	"go-designpatterns/strategy"
	"go-designpatterns/decorator"
	"go-designpatterns/observer"
	"go-designpatterns/command"
	"go-designpatterns/adapter"
)

func main() {
	singleton.Run()
	strategy.Run()
	decorator.Run()
	observer.Run()
	command.Run()
	adapter.Run()
}