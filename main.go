package main

import (
	"fmt"

	IAfter "github.com/johnatasr/GO-SOLID/i_after"
	IBefore "github.com/johnatasr/GO-SOLID/i_before"
	Liskov "github.com/johnatasr/GO-SOLID/liskov"
	OAfter "github.com/johnatasr/GO-SOLID/o_after"
	OBefore "github.com/johnatasr/GO-SOLID/o_before"
	SAfter "github.com/johnatasr/GO-SOLID/s_after"
	SBefore "github.com/johnatasr/GO-SOLID/s_before"
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

	//Old I
	lb_square := IBefore.CreateSquare(3)
	lb_cube := IBefore.CreateCube(4)

	fmt.Println(IBefore.AreaSum(lb_square, lb_cube))
	fmt.Println(IBefore.AreaVolumeSum(lb_square, lb_cube))

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
	fmt.Println()

	fmt.Println("=========Liskov Principle===========")

	human := Liskov.Human{Name: "Alex"}
	student := Liskov.Student{
		Human: Liskov.Human{Name: "Mike"},
		Grades: map[string]int{
			"Math":    8,
			"English": 9,
		},
	}
	teacher := Liskov.Teacher{
		Human:  Liskov.Human{Name: "John"},
		Degree: "CS",
		Salary: 2000,
	}

	p := Liskov.Printer{}
	p.Info(human)
	p.Info(student)
	p.Info(teacher)

	fmt.Println()
	fmt.Println("=====Interfaces implementation=====")

	ia_square1 := IAfter.CreateSquare(5)
	ia_square2 := IAfter.CreateSquare(4)
	ia_cube1 := IAfter.CreateCube(ia_square1)
	ia_cube2 := IAfter.CreateCube(ia_square2)
	fmt.Println(IAfter.AreaSum(ia_square1, ia_square2, ia_cube1, ia_cube2))
	fmt.Println(IAfter.AreaVolumeSum(ia_cube1, ia_cube2))

}
