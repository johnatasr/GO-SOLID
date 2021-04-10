package i_after

import (
	"math"
)

type IShape interface {
	Area() float64
}

type Square struct {
	Length float64
}

func CreateSquare(length float64) Square {
	return Square{Length: length}
}

func (s Square) Area() float64 {
	return math.Pow(s.Length, 2)
}

type Cube struct {
	Square
}

func CreateCube(square Square) Cube {
	return Cube{Square: square}
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Length, 3)
}

type IObject interface {
	IShape
	Volume() float64
}

func AreaSum(shapes ...IShape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}

func AreaVolumeSum(shapes ...IObject) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.Area() + s.Volume()
	}
	return sum
}

