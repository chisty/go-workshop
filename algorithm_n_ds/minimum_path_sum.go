package main

import "fmt"

func main() {
	fmt.Println("Minimum Path Sum")

	res := minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	fmt.Println("Res= ", res)
}

func minPathSum(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])

	for i := 1; i < row; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < col; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			minVal := getMin(grid[i-1][j], grid[i][j-1])
			grid[i][j] += minVal
		}
	}

	return grid[row-1][col-1]
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
