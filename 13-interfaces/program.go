package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Length float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

/* utility functions */
type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)
	PrintPerimeter(c) // (2 * pi * r)

	r := Rectangle{Length: 10, Width: 12}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)
	PrintPerimeter(r) // (2 * (length + width))
}
