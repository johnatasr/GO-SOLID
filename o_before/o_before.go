package o_before

import (
	"math"
)

type Circle struct {
	Radius float64
}

func CreateCircle(radius float64) Circle {
	return Circle{Radius: radius}
}

type Square struct {
	Length float64
}

func CreateSquare(length float64) Square {
	return Square{Length: length}
}

type Calculator struct {
}

func (calc Calculator) AreaSum(shapes ...interface{}) float64 {
	var sum float64
	for _, shape := range shapes {
		switch shape.(type) {
		case Circle:
			r := shape.(Circle).Radius
			sum += math.Pi * r * r
		case Square:
			l := shape.(Square).Length
			sum += l * l
		}
	}
	return sum
}
