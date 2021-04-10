  
package o_after

import (
	"math"
)

type IShape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func CreateCircle(radius float64) Circle {
	return Circle{Radius:radius}
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	Length float64
}

func CreateSquare(length float64) Square {
	return Square{Length:length}
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Triangule struct {
	Height float64
	Base   float64
}

func CreateTriangule(height float64, base float64) Triangule {
	return Triangule{
		Height:height,
		Base: base}
}

func (t Triangule) Area() float64 {
	return t.Base * t.Height / 2
}

type Calculator struct {
}

func (a Calculator) AreaSum(shapes ...IShape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.Area()
	}
	return sum
}

