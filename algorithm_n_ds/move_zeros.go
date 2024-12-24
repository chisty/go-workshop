package main

import "fmt"

func main() {
	fmt.Println("Move Zeros")

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}

	for i := pos; i < len(nums); i++ {
		nums[i] = 0
	}

}
