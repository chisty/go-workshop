package main

import "fmt"

func main() {
	fmt.Println("Hello")

	data := [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}

	res := arrayManipulation(5, data)
	fmt.Printf("Result= %d", res)
}

func arrayManipulation(n int32, queries [][]int32) int64 {
	data := make([]int32, n)

	var total int64 = 0
	l := len(queries)

	for i := 0; i < l; i++ {
		data[queries[i][0]-1] += queries[i][2]
		if n > (queries[i][1]) {
			data[queries[i][1]] -= queries[i][2]
		}
	}

	var m int64 = 0
	for i := 0; i < len(data); i++ {
		total += (int64)(data[i])
		if total > m {
			m = total
		}
	}

	return m

}
