package main

import "fmt"

func main() {
	var name [2]string
	name[0] = "Hatta"
	name[1] = "Sugiarto"

	fmt.Println(name[0], name[1])

	var values = [3]int{
		80, 90, 100,
	}
	fmt.Println(values)
	fmt.Print(len(name), len(values))
}
