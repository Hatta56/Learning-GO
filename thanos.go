package main

import "fmt"

func main() {
	penduduk := 1
	for i := 1; i <= 50; i++ {

		if i == 1 {
			fmt.Println("Hari ke-", i, " ada", penduduk, " Penduduk")
		} else if i%3 == 0 {
			penduduk /= 2
			fmt.Println("Hari ke-", i, " ada", penduduk, " Penduduk")
		} else {
			penduduk *= 3
			fmt.Println("Hari ke-", i, " ada", penduduk, " Penduduk")

		}
	}
}
