package main

import "fmt"

func main() {

	fmt.Println("Maximum Subarray")

	fmt.Println("Result= ", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println("Result= ", maxSubArray([]int{-2}))
	fmt.Println("Result= ", maxSubArray([]int{-2, 1, -3}))
	fmt.Println("Result= ", maxSubArray([]int{-2, 1, -3, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	sum := nums[0]
	temp := nums[0]

	for i := 1; i < len(nums); i++ {
		temp += nums[i]

		if temp < nums[i] {
			temp = nums[i]
		}

		if temp > sum {
			sum = temp
		}
	}

	return sum
}
