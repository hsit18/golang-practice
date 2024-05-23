package structsdemo

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func (p *person) updateFirstName(name string) {
	p.firstname = name
}

func (p *person) Display() {
	fmt.Println(p)
	fmt.Printf("%+v", p)
}

func Execute() {
	ravi := person{firstname: "Ravi", lastname: "Sharma"}
	ravi.Display()
	ravi.updateFirstName("Rahul")
	ravi.Display()
}
