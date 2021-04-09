package s_before

import (
	"fmt"
	"math"
)

type Square struct {
	Lenght float64
}

func CreateSquare(lenght float64) *Square {
	return &Square{Lenght: lenght}
}

func (s *Square) Area() {
	fmt.Printf("Square are is: %f\n", s.Lenght * s.Lenght)
}


type Circle struct {
	Radius float64
}

func CreateCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

func (c *Circle) Area() {
	fmt.Printf("Circle are is: %f\n", math.Pi * c.Radius * c.Radius)
}



