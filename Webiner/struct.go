package main

import "fmt"

type Hitung struct {
	angka1 int
	angka2 int
}

func (h Hitung) perhitungan() (hasil int) {
	hasil = h.angka1 + h.angka2

	return

}

func main() {
	menghitung := Hitung{
		angka1: 20,
		angka2: 30,
	}

	fmt.Println(menghitung.perhitungan())
}
