package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	ID       string `validate:"omitempty,uuid"`
	Name     string `validate:"required"`
	Username string `validate:"required,email"`
	Password string `validate:"required,gte=10"`
	Type     string `validate:"isdefault"`
}

func main() {
	user := User{
		Name:     "Ahmed Chisty",
		Username: "ahmed@gmail.com",
		Password: "1234567890",
	}

	validate := validator.New()
	error := validate.Struct(user)
	if error != nil {
		fmt.Println("Error Found: ", error)
	} else {
		fmt.Println("Valid Struct found.")
	}
}
