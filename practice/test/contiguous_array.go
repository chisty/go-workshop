package main

import "fmt"

func main() {
	fmt.Println("Contiguous Array")
	fmt.Println("Result= ", findMaxLength([]int{0, 1}))
	fmt.Println("Result= ", findMaxLength([]int{0, 1, 0}))

	fmt.Println("Result= ", findMaxLength([]int{1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1}))
	fmt.Println("Result= ", findMaxLength([]int{0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0}))
}

func findMaxLength(nums []int) int {
	store := make(map[int]int)
	total, result := 0, 0
	store[0] = -1
	for i, val := range nums {
		if val == 0 {
			val = -1
		}
		total += val
		old, ok := store[total]
		if ok {
			diff := i - old
			if diff > result {
				result = diff
				//fmt.Printf("%d-%d\n", old, i)
			}
		} else {
			store[total] = i
		}
	}
	return result
}
