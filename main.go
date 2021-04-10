package main

import (
	"fmt"

	SBefore "github.com/johnatasr/GO-SOLID/s_before"
	SAfter "github.com/johnatasr/GO-SOLID/s_after"
	OBefore "github.com/johnatasr/GO-SOLID/o_before"
	OAfter "github.com/johnatasr/GO-SOLID/o_after"
)

func main() {

	fmt.Println("======Before SOLID=======")

	// Old S
	sb_square := SBefore.CreateSquare(22)
	sb_square.Area()

	fmt.Println("===================================")

	sb_circle := SBefore.CreateCircle(10)
	sb_circle.Area()

	//Old O

	ob_square := OBefore.CreateSquare(22)
	ob_circle := OBefore.CreateCircle(10)

	calculator := OBefore.Calculator{}

	fmt.Println("Area sum:", calculator.AreaSum(ob_square, ob_circle))


	fmt.Println("============After SOLID============")
	fmt.Println("===================================")
	fmt.Println("======Single Responsabilitie=======")

	// New S
	sa_square := SAfter.NewSquare(22)
	sa_circle := SAfter.NewCircle(10)

	out := SAfter.Outputter{}
	fmt.Println(out.Text(sa_square))
	fmt.Println(out.Text(sa_circle))

	fmt.Println("============Json Return============")
	fmt.Println(out.JSON(sa_square))
	fmt.Println(out.JSON(sa_circle))

	fmt.Println("============Open Closed============")
	oa_circle := OAfter.CreateCircle(10)
	oa_square := OAfter.CreateSquare(22)
	oa_triangule := OAfter.CreateTriangule(3, 7)
	calc := OAfter.Calculator{}
	fmt.Println("area sum:", calc.AreaSum(oa_circle, oa_square, oa_triangule))

}