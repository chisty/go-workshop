package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	DOB       string   `json:"dob"`
	Address   *Address `json:address`
}

type Address struct {
	Country  string `json:"country"`
	Location string `json:"location"`
}

func main() {
	csvFile, error := os.Open("a.xlsx")
	if error != nil {
		fmt.Println("Error: ", error.Error())
	}

	reader := csv.NewReader(csvFile)

	var people []Person

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println("Read error: ", error.Error())
		}

		item := Person{
			FirstName: line[0],
			LastName:  line[1],
			DOB:       line[4],
			Address: &Address{
				Location: line[2],
				Country:  line[3],
			},
		}

		people = append(people, item)
	}

	resultJSON, error := json.Marshal(people)
	if error != nil {
		fmt.Println("Error in marshaling: ", error.Error())
	}

	fmt.Println(string(resultJSON))
}
