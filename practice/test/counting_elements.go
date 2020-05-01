package main

import "fmt"

func main() {
	fmt.Println("Counting Elements")
	fmt.Println("2= ", countElements([]int{1, 2, 3}))
	fmt.Println("0= ", countElements([]int{1, 1, 3, 3, 5, 5, 7, 7}))
	fmt.Println("3= ", countElements([]int{1, 3, 2, 3, 5, 0}))
	fmt.Println("2= ", countElements([]int{1, 1, 2, 2}))
	fmt.Println("2= ", countElements([]int{2, 1, 1}))
	fmt.Println("4= ", countElements([]int{1, 3, 2, 3, 5, 2, 0}))
	fmt.Println("4= ", countElements([]int{1, 3, 2, 5, 2, 0}))

}

func countElements(arr []int) int {
	total := 0
	var data [1001]int

	for _, val := range arr {
		if data[val] == 0 {
			if val+1 < 1001 && data[val+1] > 0 {
				total++
			}
			if val-1 > -1 && data[val-1] > 0 {
				total += data[val-1]
			}
		} else {
			if val+1 < 1001 && data[val+1] > 0 {
				total++
			}
		}
		data[val]++
	}

	return total
}
