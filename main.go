package main

import (
	"fmt"

	SBefore "github.com/johnatasr/GO-SOLID/s_before"
)

func main(){

	// Old S
	square := SBefore.CreateSquare(22)
	square.Area()

	fmt.Println("===================================")

	circle := SBefore.CreateCircle(10)
	circle.Area()



}