package main

import (
	"fmt"
)

func main() {
	var nilai32 int32 = 1000
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai64)

	var name = "Hatta"
	var e = name[0]
	var eString string = string(e)

	fmt.Println(eString)
	fmt.Println(e)
}
