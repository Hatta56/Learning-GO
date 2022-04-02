package main

import "fmt"

type Addres struct {
	City, Province, Country string
}

func changeCountry(addres *Addres) {
	addres.Country = "China"
}
func main() {
	addres := Addres{"Lebak", "Banten", "Indonesia"}
	addres1 := &addres

	addres1.City = "Jakarta"
	*addres1 = Addres{"Malang", "Banten", "Indonesia"}
	fmt.Println(addres)
	fmt.Println(addres1)

	var addres4 *Addres = new(Addres)

	addres4.City = "Jakarta"
	fmt.Println(addres4)

	alamat := Addres{
		City:     "Wuhan",
		Province: "Lombok",
		Country:  "",
	}
	alamatPointer := &alamat
	changeCountry(alamatPointer)
	fmt.Println(alamat)
}
