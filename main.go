package main

import (
	"go-designpatterns/singleton"
	"go-designpatterns/strategy"
	"go-designpatterns/decorator"
)

func main() {
	singleton.Run()
	strategy.Run()
	decorator.Run()
}