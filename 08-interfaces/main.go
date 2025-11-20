package main

import (
	"fmt"
	"math"
)

// 1. Defining an Interface
// Any type that has an Area() and Perimeter() method satisfies this interface.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. Implementing the Interface (Rectangle)
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// 3. Implementing the Interface (Circle)
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 4. Polymorphism
// This function takes ANY Shape. It doesn't care if it's a Rectangle or Circle.
func printShapeInfo(s Shape) {
	fmt.Printf("Type: %T\n", s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	fmt.Println("---")
}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 5}

	// Both r and c are Shapes!
	printShapeInfo(r)
	printShapeInfo(c)
}
