package main

import "fmt"

func main() {

	person := map[string]string{
		"name":     "Hatta",
		"LastName": " Sugiarto",
	}

	for index, value := range person {
		fmt.Println(index, value)
	}

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Hatta Sugiarto"
	book["soalah"] = "salah euy"

	fmt.Println(book)

	delete(book, "soalah")
	fmt.Println(book)
}
