package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string `json:"firstName"` // key akan menjadi "firstName"
	BirthYear int    `json:"birtYear"`  // key akan menjadi "birthYear"
	Email     string // key akan menjadi "Email"
	Address   string `json:"address,omitempty"`
	Hobby     string `json:"-"`
}

func main() {
	byteJSONData := []byte(`{
		"FirstName": "Rogu",
		"BirthYear": 2000,
		"Email":     "example@ruangguru.com",
		"Address":   "bb",
		"Hobby":     "belajar",
		"age":20
	}`)

	u := User{}
	err := json.Unmarshal(byteJSONData, &u)
	fmt.Println(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("first name: ", u.FirstName)
	fmt.Println("birth year: ", u.BirthYear)
}
