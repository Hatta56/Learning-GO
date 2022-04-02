package main

import (
	"fmt"
)

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()

	switch value := result.(type) {
	case string:
		fmt.Println("Ini adalah string", value)
	case int:
		fmt.Println("ini adalah Int", value)
	default:
		fmt.Println("type not found")
	}

}
