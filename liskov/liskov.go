package main

import "fmt"

type Human struct {
	Name string
}

func (h Human) getName() string {
	return h.Name
}

type Teacher struct {
	Human
	Degree string
	Salary float64
}

type Student struct {
	Human
	Grades map[string]int
}

type Person interface {
	getName() string
}

type Printer struct {
}

func (Printer) info(p Person) {
	fmt.Println("Name: ", p.getName())
}

func main() {
	h := Human{Name: "Alex"}
	s := Student{
		Human: Human{Name: "Mike"},
		Grades: map[string]int{
			"Math":    8,
			"English": 9,
		},
	}
	t := Teacher{
		Human:  Human{Name: "John"},
		Degree: "CS",
		Salary: 2000,
	}

	p := Printer{}
	p.info(h)
	p.info(s)
	p.info(t)
}
