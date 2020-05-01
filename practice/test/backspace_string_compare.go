package main

import (
	"fmt"
)

func main() {
	fmt.Println("Counting Elements")
	fmt.Println("res= ", backspaceCompare("ab#c", "ad#c"))
	fmt.Println("res= ", backspaceCompare("ab#c", "ad##ac"))
	fmt.Println("res= ", backspaceCompare("ab##", "c#d#"))
	fmt.Println("res= ", backspaceCompare("a##c", "#a#c"))
	fmt.Println("res= ", backspaceCompare("a#c", "b"))
	fmt.Println("res= ", backspaceCompare("bbbb###abcdef######", "a##b"))
	fmt.Println("res= ", backspaceCompare("###ba#bb##", "b"))
	fmt.Println("res= ", backspaceCompare("bxj##tw", "bxo#j##tw"))

}

func backspaceCompare(S string, T string) bool {
	return false
}

func backspaceCompare_old(S string, T string) bool {
	for i, j := len(S)-1, len(T)-1; i > -1 || j > -1; i, j = i-1, j-1 {
		countS, countT := 0, 0
		for i > -1 && (S[i] == '#' || countS > 0) {
			if S[i] == '#' {
				i, countS = i-1, countS+1
			} else {
				i, countS = i-1, countS-1
			}
		}

		for j > -1 && (T[j] == '#' || countT > 0) {
			if T[j] == '#' {
				j, countT = j-1, countT+1
			} else {
				j, countT = j-1, countT-1
			}
		}

		if i < 0 && j < 0 {
			return true
		}
		if i < 0 || j < 0 {
			return false
		}

		if S[i] != T[j] {
			return false
		}
	}
	return true
}
