package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var newShape Shape

	newShape = Rectangle{width: 5, height: 4}
	fmt.Println("Area of rectangle: ", newShape.Area())
	fmt.Println("Perimeter of rectangle: ", newShape.Perimeter())

	newShape = Circle{radius: 4}
	fmt.Println("Area of circle: ", newShape.Area())
	fmt.Println("Perimeter of circle: ", newShape.Perimeter())
}
