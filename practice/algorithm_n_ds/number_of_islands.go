package main

import "fmt"

func main() {
	fmt.Println("Number of Islands")
	doTest([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}})
	doTest([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}})
	doTest([][]byte{{'1', '1', '1'}, {'1', '1', '1'}, {'1', '1', '1'}})
}

func doTest(grid [][]byte) {
	fmt.Println(grid)
	result := numIslands(grid)
	fmt.Println(result)
}

func numIslands(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}

	var store []*cell
	col, count := len(grid[0]), 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				grid[i][j], count = '0', count+1
				store = append(store, newCell(i, j))

				for len(store) > 0 {
					cur := store[0]

					if cur.row > 0 && grid[cur.row-1][cur.col] == '1' {
						grid[cur.row-1][cur.col], store = '0', append(store, newCell(cur.row-1, cur.col))
					}
					if cur.col > 0 && grid[cur.row][cur.col-1] == '1' {
						grid[cur.row][cur.col-1], store = '0', append(store, newCell(cur.row, cur.col-1))
					}
					if cur.row+1 < row && grid[cur.row+1][cur.col] == '1' {
						grid[cur.row+1][cur.col], store = '0', append(store, newCell(cur.row+1, cur.col))
					}
					if cur.col+1 < col && grid[cur.row][cur.col+1] == '1' {
						grid[cur.row][cur.col+1], store = '0', append(store, newCell(cur.row, cur.col+1))
					}
					store = store[1:]
				}
			}
		}
	}

	return count
}

//Still good since it is using a fixed slice.
func numIslandsOld(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}

	col, count, si := len(grid[0]), 0, 0
	store := make([]*cell, row*col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				grid[i][j], store[0], count, si = '0', newCell(i, j), count+1, 1
				for temp := 0; temp < si; temp++ {
					cur := store[temp]

					if cur.row > 0 && grid[cur.row-1][cur.col] == '1' {
						grid[cur.row-1][cur.col], store[si], si = '0', newCell(cur.row-1, cur.col), si+1
					}
					if cur.col > 0 && grid[cur.row][cur.col-1] == '1' {
						grid[cur.row][cur.col-1], store[si], si = '0', newCell(cur.row, cur.col-1), si+1
					}
					if cur.row+1 < row && grid[cur.row+1][cur.col] == '1' {
						grid[cur.row+1][cur.col], store[si], si = '0', newCell(cur.row+1, cur.col), si+1
					}
					if cur.col+1 < col && grid[cur.row][cur.col+1] == '1' {
						grid[cur.row][cur.col+1], store[si], si = '0', newCell(cur.row, cur.col+1), si+1
					}
				}
			}
		}
	}

	return count
}

type cell struct {
	row int
	col int
}

func newCell(row, col int) *cell {
	return &cell{row, col}
}
