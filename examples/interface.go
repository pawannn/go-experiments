//An interface in Go is a type that defines a set of method signatures, and any type that implements those methods automatically satisfies the interface.

package main

import "fmt"

type shape interface {
	Area() float64
	Perimeter() float64
}

type rectange struct {
	length  float64
	breadth float64
}

func (r rectange) Area() float64 {
	return r.length * r.breadth
}

func (r rectange) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

type square struct {
	side float64
}

func (s square) Area() float64 {
	return s.side * s.side
}

func (s square) Perimeter() float64 {
	return 4 * s.side
}

func CalcArea(s shape) float64 {
	return s.Area()
}

func CalcPerimeter(s shape) float64 {
	return s.Perimeter()
}

func calc() {
	square := square{
		side: 4,
	}

	rectangle := rectange{
		length:  2,
		breadth: 4,
	}

	squareArea := CalcArea(square)
	squarePerimeter := CalcPerimeter(square)

	rectangleArea := CalcArea(rectangle)
	rectanglePerimeter := CalcPerimeter(rectangle)

	fmt.Println("square area : ", squareArea, " square perimeter : ", squarePerimeter)
	fmt.Println("rectangle area : ", rectangleArea, " rectangle perimeter : ", rectanglePerimeter)
}
