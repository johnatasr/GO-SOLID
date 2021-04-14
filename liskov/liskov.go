package liskov

import "fmt"

type Human struct {
	Name string
}

func (h Human) GetName() string {
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
	GetName() string
}

type Printer struct {
}

func (Printer) Info(p Person) {
	fmt.Println("Name: ", p.GetName())
}

func main() {

}
