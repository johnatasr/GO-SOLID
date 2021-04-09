package main

import (
	"fmt"

	SBefore "github.com/johnatasr/GO-SOLID/s_before"
	SAfter "github.com/johnatasr/GO-SOLID/s_after"
)

func main(){

	// Old S
	square := SBefore.CreateSquare(22)
	square.Area()

	fmt.Println("===================================")

	circle := SBefore.CreateCircle(10)
	circle.Area()


	// New S
	square := SAfter.NewSquare(22)
	circle := SAfter.NewCircle(10)

	out := SAfter.Outputter{}
	out.Text(square)
	out.Text(circle)


}