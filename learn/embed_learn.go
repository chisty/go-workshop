package main

import (
	"fmt"
	"sort"
)

type customReverse struct {
	sort.Interface
}

func (r customReverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

func RunEmbed1() {
	data := []int{4, 5, 2, 8, 1, 9, 3}
	sort.Sort(sort.IntSlice(data))
	fmt.Println("After library Sort: ", data)

	rev := &customReverse{sort.IntSlice(data)}
	sort.Sort(rev)
	fmt.Println("After Custom Sort: ", data)
}
