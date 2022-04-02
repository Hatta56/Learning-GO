package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	slice := []string{}
	var arrays []string
	for i := 1; i <= 1111; i++ {
		slice = append(slice, strconv.Itoa(i))
		arrays = strings.Split(slice[i-1], "")

		for _, v := range arrays {
			fmt.Printf("Karakter ke %v adalah %v\n", len(slice), v)
		}
	}
}
