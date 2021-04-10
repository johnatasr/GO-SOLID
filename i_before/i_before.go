package i_before

import (
	"math"
)

type IShape interface {
	Area() float64
	Volume() float64
}

type Square struct {
	Length float64
}

func CreateSquare(length float64) Square {
	return Square{Length: length}
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

// square is a flat surface -> does not have volume
func (s Square) Volume() float64 {
	return 0
}

type Cube struct {
	Length float64
}

func CreateCube(length float64) Cube {
	return Cube{Length: length}
}

func (c Cube) Area() float64 {
	return math.Pow(c.Length, 2)
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Length, 3)
}

func AreaSum(shapes ...IShape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}

func AreaVolumeSum(shapes ...IShape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.Area() + s.Volume()
	}
	return sum
}

