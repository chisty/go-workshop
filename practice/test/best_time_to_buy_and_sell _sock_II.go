package main

import "fmt"

func main() {
	fmt.Println("Best Time To Buy & Sell Stock II")

	fmt.Println("Result= ", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("Result= ", maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println("Result= ", maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println("Result= ", maxProfit([]int{7, 1, 3, 6, 4, 3, 1}))
	fmt.Println("Result= ", maxProfit([]int{}))

}

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

//old Code
func maxProfit_old(prices []int) int {
	result, low, high := 0, 0, 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[high] {
			result += prices[high] - prices[low]
			low, high = i, i
		} else {
			high++
		}
	}

	if high < len(prices) && low < len(prices) {
		result += prices[high] - prices[low]
	}

	return result

}
