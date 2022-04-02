package main

import "fmt"

func Test(i int) interface{} {
	if i == 2 {
		return 1
	} else if i == 3 {
		return true
	} else {
		return "Hatta"
	}
}

func main() {
	var Call interface{} = Test(1)

	fmt.Println(Call)
}
