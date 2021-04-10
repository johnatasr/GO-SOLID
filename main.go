package main

import (
	"fmt"

	SBefore "github.com/johnatasr/GO-SOLID/s_before"
	SAfter "github.com/johnatasr/GO-SOLID/s_after"
)

func main(){

	// Old S
	b_square := SBefore.CreateSquare(22)
	b_square.Area()

	fmt.Println("===================================")

	b_circle := SBefore.CreateCircle(10)
	b_circle.Area()

	fmt.Println("===================================")
	fmt.Println("======Single Responsabilitie=======")

	// New S
	a_square := SAfter.NewSquare(22)
	a_circle := SAfter.NewCircle(10)

	out := SAfter.Outputter{}
	fmt.Println(out.Text(a_square))
	fmt.Println(out.Text(a_circle))

	fmt.Println("============Json Return============")
	fmt.Println(out.JSON(a_square))
	fmt.Println(out.JSON(a_circle))
}