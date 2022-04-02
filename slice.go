package main

import "fmt"

func main() {
	var month = [...]string{
		"januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desembre",
	}

	var slice = month[3:5]

	fmt.Println(len(slice), cap(slice))

	var slice2 = month[10:]

	var slice3 = append(slice2, "Hatta")
	slice3[2] = "Siswandi"
	fmt.Println(slice3)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Hatta"
	newSlice[1] = "Sugiarto"

	fmt.Println(len(newSlice))
	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))

	toSlice := make([]string, len(newSlice), cap(newSlice))

	copy(toSlice, newSlice)
	fmt.Println(toSlice)

	//perbedaan array dan slice

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
