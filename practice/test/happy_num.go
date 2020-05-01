package main

import (
	"fmt"
)

func main() {
	fmt.Println("Happy Number")
	fmt.Println("Happy= ", isHappy(19))
	for i := 0; i < 20; i++ {
		fmt.Printf("%d is Happy= %t\n", i, isHappy(i))
	}
}

func isHappy(n int) bool {
	store := make(map[int]bool)
	return recursiveCheck(n, store)
}

func recursiveCheck(n int, store map[int]bool) bool {
	if n == 1 {
		return true
	}

	if _, ok := store[n]; ok {
		return false
	}
	store[n] = false

	return recursiveCheck(getSum(n), store)
}

// func isHappy(n int) bool {
// 	if n < 1 {
// 		return false
// 	}

// 	trueList := []bool{false, true, false, false, false, false, false, true, false, false}
// 	store := make(map[int]bool)
// 	for i := 1; i < 10; i++ {
// 		store[i] = trueList[i]
// 		store[i*i] = trueList[i]
// 	}

// 	for {
// 		value, ok := store[n]
// 		if ok {
// 			return value
// 		}
// 		store[n] = false
// 		n = getSum(n)
// 	}

// 	return false
// }

func getSum(n int) int {
	sum := 0
	for n > 0 {
		temp := n % 10
		sum += temp * temp
		n /= 10
	}
	return sum
}
