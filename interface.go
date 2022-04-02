package main

import "fmt"

type hasName interface {
	GetName() string
}

func sayHello(s hasName) {
	fmt.Print("Hello", s.GetName())
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}
func main() {

	People := Person{
		Name: "Hatta",
	}

	sayHello(People)
}
