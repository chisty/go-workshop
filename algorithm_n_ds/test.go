package main

import "fmt"

func main() {
	fmt.Println("Test and R&D Code")
	doTest([]int{10, 20, 30, 40, 50, 60})
}

func doTest(data []int) {
	printArray(data)
	update(data)
	printArray(data)
}

func update(data []int) {
	data[0] = 100
}

func printArray(data []int) {
	for i := 0; i < len(data); i++ {
		fmt.Printf("%d	", data[i])
	}
	fmt.Println()
}
