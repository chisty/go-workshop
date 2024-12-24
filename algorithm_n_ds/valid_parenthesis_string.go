package main

import "fmt"

func main() {
	fmt.Println("Valid Parenthesis String")
	doTest("(())(())(((()*()()()))()((()()(*()())))(((*)()")
	doTest("(((()*())))((()(((()(()))()**(*)())))())()()*")
	doTest("(()())")
	doTest("")
	doTest("(())")
	doTest("(***")
	doTest("((**")

	doTest("((***))(((")
	doTest("((***)***)(((")

	doTest("(*))")
	doTest("*(*)**))")
	doTest("(*)")
	doTest("(())**(()))**")
	doTest("(((***))***())")
	doTest("(((***)***())")
	doTest("(((**)*()))")
	doTest("(((**)*()")

}

func doTest(s string) {
	fmt.Printf("%s= ", s)
	res := checkValidString(s)
	fmt.Println(res)
}

func checkValidString(s string) bool {
	input, l, count := []rune(s), len(s), 0

	for i := 0; i < l; i++ {
		if input[i] == '(' || input[i] == '*' {
			count++
		} else {
			count--
		}
		if count < 0 {
			return false
		}
	}

	count = 0
	for i := l - 1; i > -1; i-- {
		if input[i] == ')' || input[i] == '*' {
			count++
		} else {
			count--
		}
		if count < 0 {
			return false
		}
	}

	return true
}

func checkValidString_not_accepted(s string) bool {
	in, l, star, left, right, flag := []rune(s), len(s), 0, 0, 0, false

	for i := 0; i < l; i++ {
		if in[i] == '*' {
			star++
		} else if in[i] == '(' {
			if flag {
				if left == right {
					left, right, star, flag = 0, 0, 0, false
				} else if left > right {
					left = left - right
					right, flag = 0, false
				} else {
					star = star + left - right
					if star < 0 {
						return false
					}
					left, right, flag = 0, 0, false
				}
			}
			left++
		} else if in[i] == ')' {
			flag = true
			right++
		}
	}

	diff := getDifference(left, right)
	fmt.Printf("left= %d, right= %d, star= %d", left, right, star)
	if diff == 0 || diff <= star {
		return true
	}

	return false
}

func getDifference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// type item struct {
// 	left int
// 	star int
// }

// type stack []item

// func newStack(int size) *stack {
// 	stack := make([]item, size)
// 	return &stack
// }

// func (s *stack) push(item r) {
// 	s = append(s, r)
// }

// func (s *stack) pop() item {
// 	l := len(stack)
// 	r := stack[l-1]
// 	stack = stack[:l]
// 	return r
// }

// func (s *stack) top() item {
// 	return s[len(s)-1]
// }

// func (s *stack) isEmpty() bool {
// 	return len(s) == 0
// }
