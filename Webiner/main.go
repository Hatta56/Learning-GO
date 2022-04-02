package main

import "fmt"

func twice(f func(x int) int, number int) int {
	return f(f(number))
}

func main() {
	inc := func(number int) int {
		return number + 1
	}
	dec := func(number int) int {
		return number - 1
	}
	fmt.Println(twice(inc, 5)) // 7
	fmt.Println(twice(dec, 5)) // 3
}
