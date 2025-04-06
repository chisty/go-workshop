package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello Concurrency!")

	var once sync.Once
	once.Do(func() {
		fmt.Println("First")
	})
	once.Do(func() {
		fmt.Println("Second")
	})
}
