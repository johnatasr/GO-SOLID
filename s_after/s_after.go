package s_after

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type IShape interface {
	Area() float64
	Name() string
}

type Square struct {
	Length float64
}

func NewSquare(length float64) *Square {
	return &Square{Length: length}
}

func (s Square) Name() string {
	return "Square"
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

func (c Circle) Name() string {
	return "Circle"
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Outputter struct {}

func (o Outputter) Text(s IShape) string {
	return fmt.Sprintf("Area of the %s: %f", s.Name(), s.Area())
}

func (o Outputter) JSON(s IShape) string {
	res := struct {
		Name string `json:"shape"`
		Area float64 `json:"area"`
	}{
		Name: s.Name(),
		Area: s.Area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}