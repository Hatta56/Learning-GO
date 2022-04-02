package main

import "fmt"

func main() {
	name := "tya"
	Length := len(name)
	if name == "Hatta" {
		fmt.Println("Sugiarto")
	} else if name == "Siswandi" {
		fmt.Println("Hallo")
	} else if Length > 5 {
		fmt.Println(name + "Halle")
	} else {
		fmt.Println("Siapa ?")
	}
}
