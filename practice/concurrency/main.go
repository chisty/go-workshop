package main

import (
	"fmt"
	"math/rand"
)

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})

	go func() {
		defer close(valueStream)

		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func primeFinder(n int) int {
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return 0
		}
	}
	return n
}

func main() {
	fmt.Println("Hello")

	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} { return rand.Intn(50000000) }

	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}

}
