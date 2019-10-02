package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   struct {
		Street  string
		City    string
		Country string
	}
}

func main() {
	mapPerson := make(map[string]interface{})
	mapPerson["FirstName"] = "Z Ahmed"
	mapPerson["LastName"] = "Chisty"
	mapPerson["Age"] = 31

	address := make(map[string]interface{})
	address["Street"] = "Free School Street"
	address["City"] = "Dhaka"
	address["Country"] = "Bangladesh"

	mapPerson["Address"] = address

	fmt.Println("Map:= ", mapPerson)

	var person Person

	err := mapstructure.Decode(mapPerson, &person)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Decoded: ")
	fmt.Println(person)
}
