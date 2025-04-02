package main

import (
	"go-designpatterns/singleton"
	"go-designpatterns/strategy"
	"go-designpatterns/decorator"
	"go-designpatterns/observer"
	"go-designpatterns/command"
	"go-designpatterns/adapter"
	"go-designpatterns/facade"
	"go-designpatterns/template"
	"go-designpatterns/iterator"
)

func main() {
	singleton.Run()
	strategy.Run()
	decorator.Run()
	observer.Run()
	command.Run()
	adapter.Run()
	facade.Run()
	template.Run()
	iterator.Run()
}