// Composition in Go is the practice of building complex types by embedding other types (structs or interfaces) inside a struct to reuse their behavior instead of using inheritance.

package main

import "fmt"

type EngineInterface interface {
	start()
	stop()
}

type CarEngine struct{}

func (e CarEngine) start() {
	fmt.Println("Car Engine started")
}

func (e CarEngine) stop() {
	fmt.Println("Car Engine stopped")
}

type TruckEngine struct{}

func (e TruckEngine) start() {
	fmt.Println("Car Engine started")
}

func (e TruckEngine) stop() {
	fmt.Println("Car Engine stopped")
}

type Transmission struct{}

func (t Transmission) shiftup() {
	fmt.Println("shifting up")
}

func (t Transmission) shiftDown() {
	fmt.Println("shifting down")
}

type steeringWheel struct{}

func (s steeringWheel) TurnRight() {
	fmt.Println("turning right")
}

func (s steeringWheel) TurnLeft() {
	fmt.Println("turning left")
}

type convertible struct {
	EngineInterface
	Transmission
	steeringWheel
}

func (c convertible) Convert() {
	fmt.Println("converting...")
}

type Truck struct {
	EngineInterface
	Transmission
	steeringWheel
}

func (t Truck) FourWheelDriver() {
	fmt.Println("Toogling 4WD...")
}
