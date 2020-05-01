package main

import "fmt"

func main() {
	fmt.Println("Hello")
	x := []int{2, 2, 1}
	y := []int{4, 1, 2, 1, 2}

	fmt.Println(singleNumber1(x))
	fmt.Println(singleNumber1(y))

}

func singleNumber1(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}
