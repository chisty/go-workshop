package main

import "fmt"

type user struct {
	name string
}

func main() {
	u := &user{name: "Chisty"}

	fmt.Println("Before: ", u)
	change(u)
	fmt.Println("After: ", u)
	// reader := make(chan int)
	// operator := make(chan int)

	// go read(reader)
	// go operate(reader, operator)

	// for value := range operator {
	// 	fmt.Println(value)
	// }
}

func change(u *user) {
	//u.name = "Ahmed Chisty"
	*u = user{name: "Ahmed Chisty"}
}

func read(r chan<- int) {
	for i := 1; i < 11; i++ {
		r <- i
	}
	close(r)
}

func operate(r <-chan int, o chan<- int) {
	for value := range r {
		o <- value * value
	}
	close(o)
}

func printer(r chan int) {
	for value := range r {
		fmt.Println(value)
	}
}

type line struct {
	value []int
}

func (data line) GetNewSlice() line {
	newList := make([]int, len(data.value))
	for i, value := range data.value {
		newList[i] = value
	}
	return line{value: newList}
}

func (data line) Update() {
	for i := 0; i < len(data.value); i++ {
		data.value[i] += 10
	}
}
