package main

import "fmt"

func main() {
	fmt.Println("Product of Array Except Self")

	fmt.Println("Result= ", productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println("Result= ", productExceptSelf([]int{1, 1}))
	fmt.Println("Result= ", productExceptSelf([]int{1, -1}))
	fmt.Println("Result= ", productExceptSelf([]int{0, 0}))
	fmt.Println("Result= ", productExceptSelf([]int{0, 1, 2, 3}))
	fmt.Println("Result= ", productExceptSelf([]int{1, 2, 3, 0}))
	fmt.Println("Result= ", productExceptSelf([]int{5, 2, 0, 10, 5}))
}

//using 1 slice
func productExceptSelf(nums []int) []int {
	l := len(nums)
	last := make([]int, l)
	last[l-1] = nums[l-1]
	for i := l - 2; i > -1; i-- {
		last[i] = nums[i] * last[i+1]
	}

	first := nums[0]
	nums[0] = last[1]

	for i := 1; i < l; i++ {
		if i == l-1 {
			nums[i] = first
			continue
		}
		nums[i], first = first*last[i+1], first*nums[i]
	}

	return nums
}

//using 2 slice
func productExceptSelf_old(nums []int) []int {
	l := len(nums)
	first, last := make([]int, l), make([]int, l)
	first[0], last[l-1] = nums[0], nums[l-1]
	for i := 1; i < l; i++ {
		first[i] = nums[i] * first[i-1]
	}
	for i := l - 2; i > -1; i-- {
		last[i] = nums[i] * last[i+1]
	}

	nums[0] = last[1]
	nums[l-1] = first[l-2]

	for i := 1; i < l-1; i++ {
		nums[i] = first[i-1] * last[i+1]
	}

	return nums
}
