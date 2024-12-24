package main

import "fmt"

func main() {
	fmt.Println("Leftmost Column with at Least a One")
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dim := binaryMatrix.Dimensions()
	row, col, res := (dim[0] - 1), (dim[1] - 1), -1

	for col := col - 1; col > -1 && row > -1; {
		val := binaryMatrix.Get(row, col)
		if val == 1 {
			res = col
			col--
		} else {
			row--
		}
	}
	return res
}

/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */
