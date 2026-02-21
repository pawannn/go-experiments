// Abstraction is the concept of hiding implementation details and exposing only the necessary behavior through a simple interface.

package main

import "fmt"

type vendingMachine interface {
	GetDrink(money int, brand string) string
}

type Application struct {
	vm vendingMachine
}

func (a Application) Run() {
	myDrink := a.vm.GetDrink(100, "cola")
	fmt.Println(myDrink)
}

func NewApplication(vm vendingMachine) *Application {
	return &Application{vm: vm}
}
