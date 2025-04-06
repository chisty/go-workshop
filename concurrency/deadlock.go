package main

import (
	"fmt"
	"sync"
	"time"
)

type Value struct {
	mu    sync.Mutex
	value int
}

func Deadlock() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *Value) {
		defer wg.Done()

		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)

		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	wg.Add(2)

	var a, b Value

	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}

func LiveLock() {
	var wg sync.WaitGroup
	spot := make(chan struct{}, 1)

	spot <- struct{}{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-spot:
				fmt.Println("Person 1: I took the spot")
				time.Sleep(100 * time.Millisecond)
				fmt.Println("Peron 1: I will release the spot!")
				spot <- struct{}{}

			default:
				fmt.Println("Person 1: Spot's not available! I wil retry again.")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for {
			select {
			case <-spot:
				fmt.Println("Person 2: I took the spot")
				time.Sleep(100 * time.Millisecond)
				fmt.Println("Peron 2: I will release the spot!")
				spot <- struct{}{}

			default:
				fmt.Println("Person 2: Spot's not available! I wil retry again.")
				time.Sleep(50 * time.Millisecond)
			}
		}

	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Main: Stopping container")
}
