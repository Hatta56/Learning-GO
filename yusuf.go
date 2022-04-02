package main

import "fmt"

func getUnik(nilaiLama, nilaiBaru int) int {
	res := (nilaiLama * nilaiBaru) % 100007
	return res
}

func main() {
	nilaiLama := 123
	nilaiBaru := 1
	for i := 0; i <= 10; i++ {
		res := getUnik(nilaiLama, nilaiBaru)
		nilaiBaru = res
		fmt.Println(res)
	}
}
