package main

import (
	"fmt"
)

// Define an interface for shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct for rectangle
type Rectangle struct {
	length float64
	width  float64
}

// Implement the Area method for Rectangle struct
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Implement the Perimeter method for Rectangle struct
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	var length, width float64

	fmt.Print("Enter the length of the rectangle: ")
	fmt.Scanln(&length)

	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scanln(&width)

	rectangle := Rectangle{length, width}

	fmt.Printf("Area: %.2f\n", rectangle.Area())
	fmt.Printf("Perimeter: %.2f\n", rectangle.Perimeter())
}
