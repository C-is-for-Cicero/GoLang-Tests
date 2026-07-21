package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{radius: 2}

	fmt.Println("Rectangle area:", r.Area())
	fmt.Println("Circle area:", c.Area())

	shapes := []Shape{r, c}
	fmt.Println("Total area:", TotalArea(shapes))
}
