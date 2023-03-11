package main

import (
	"fmt"
	"math"
)

const Pi float64 = math.Pi

type Rectangle struct {
	name          string
	width, height float64
}

type Circle struct {
	name   string
	radius float64
}

func (c Circle) Area() float64 {
	return Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Type() string {
	return c.name
}

func (r Rectangle) Type() string {
	return r.name
}

func main() {
	r := Rectangle{
		name:   "Прямоугольник",
		width:  12,
		height: 8,
	}
	c := Circle{
		name:   "Окружность",
		radius: 6,
	}

	fmt.Printf("Type: %v\n", r.Type())
	fmt.Printf("Area: %v\n", r.Area())

	fmt.Printf("Type: %v\n", c.Type())
	fmt.Printf("Area: %v\n", c.Area())
}
