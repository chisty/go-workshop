package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Printing Fibonacci numbers..")

	reader := bufio.NewReader(os.Stdin)
	for {
		generator := calculateFibonacci()
		input, _ := reader.ReadString('\n')
		value, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
		var result int
		for i := 0; i < value; i++ {
			result = generator()
		}
		fmt.Printf("Fibonacci number %d is: %d\n", value, result)
	}
}

func calculateFibonacci() func() int {
	f1 := 0
	f2 := 1

	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}
