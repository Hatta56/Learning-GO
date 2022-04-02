package main

import "fmt"

type Man struct {
	Name string
}

func (m *Man) GetMarried() {
	m.Name = "Mr. " + m.Name
}

func main() {
	man := Man{"Hatta"}
	man.GetMarried()
	fmt.Println(man.Name)
}
