package main

import "fmt"

func main() {
	fmt.Println("Perform String Shifts")
	stringShift("abc", [][]int{{0, 4}, {1, 5}})
	// stringShift("abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}})
}

func stringShift(s string, shift [][]int) string {
	r, l, pos := []rune(s), len(s), 0

	for _, val := range shift {
		val[1] = val[1] % l
		if val[0] == 1 {
			pos = l - val[1]
			last := r[:pos]
			r = append(r[pos:], last...)
		} else {
			first := r[:val[1]]
			r = append(r[val[1]:], first...)
		}
	}

	return string(r)
}
