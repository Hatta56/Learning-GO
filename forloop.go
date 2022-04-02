package main

import "fmt"

func main() {
	// for counter := 1; counter < 10; counter++ {
	// 	fmt.Println("INI", counter)
	// }

	slice := []string{"Hatta", "Sugiarto", "Mustofa"}
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }
	for index, val := range slice {
		fmt.Println("Index", index, "=", val)
	}
}
