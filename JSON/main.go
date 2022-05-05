package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	//	DecodeJsonToObject()
	//DecodeJsonArrayToObject()
	DecodeJSONtoMapStringInterface()
	DecodeJSONtoInterface()
	EncodeToJson()

}

func DecodeJsonToObject() {
	jsonString := `{"name":"John","age":30,"city":"New York"}`
	jsonData := []byte(jsonString)

	var P Person
	err := json.Unmarshal(jsonData, &P)
	if err != nil {
		panic(err)
	}
	fmt.Println(P)
}

func DecodeJsonArrayToObject() {
	jsonString := `[
		{"name":"John","age":30,"city":"New York"},
		{"name":"Tom","age":25,"city":"Paris"},
		{"name":"Bob","age":20,"city":"London"}
	]`
	jsonData := []byte(jsonString)

	var P []Person
	err := json.Unmarshal(jsonData, &P)
	if err != nil {
		panic(err)
	}
	for _, v := range P {
		fmt.Println("Name Person :", v.Name)
		fmt.Println("Age Person :", v.Age)
		fmt.Println("City Person :", v.City)
	}
}

func DecodeJSONtoMapStringInterface() {
	jsonString := `{"name":"John","age":30,"city":"New York"}`
	jsonData := []byte(jsonString)

	var P map[string]interface{}
	err := json.Unmarshal(jsonData, &P)
	if err != nil {
		panic(err)
	}

	fmt.Println(P["name"])
}
func DecodeJSONtoInterface() {
	jsonString := `{"name":"John","age":30,"city":"New York"}`
	jsonData := []byte(jsonString)

	var P interface{}
	err := json.Unmarshal(jsonData, &P)
	if err != nil {
		panic(err)
	}

	decodePerson := P.(map[string]interface{})
	fmt.Println(decodePerson["name"])
}
func EncodeToJson() {
	P := Person{
		Name: "John",
		Age:  30,
		City: "New York",
	}
	jsonData, err := json.Marshal(P)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
