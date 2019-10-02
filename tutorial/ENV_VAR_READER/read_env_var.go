package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Read Environment Varibale")
	// fmt.Println("GoPath= ", os.Getenv("GOPATH"))

	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
