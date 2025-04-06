package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func PipelineWithTake() {
	generator := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		stream := make(chan interface{})

		go func() {
			defer fmt.Println("exiting generator.")
			defer close(stream)

			for {
				select {
				case <-done:
					return
				case stream <- fn():
				}
			}
		}()
		return stream
	}

	take := func(done <-chan interface{}, inStream <-chan interface{}, num int) <-chan interface{} {
		stream := make(chan interface{})

		go func() {
			defer fmt.Println("exiting take.")
			defer close(stream)

			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case stream <- <-inStream:
				}
			}
		}()
		return stream
	}

	randFunc := func() interface{} {
		return rand.Intn(20)
	}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, generator(done, randFunc), 5) {
		fmt.Println(num)
	}
}

func PipelineOne() {
	generator := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		stream := make(chan interface{})

		go func() {
			defer fmt.Println("exiting generator")
			defer close(stream)

			for {
				select {
				case <-done:
					return
				case stream <- fn():
				}
			}
		}()

		return stream
	}

	multiplier := func(done <-chan interface{}, inStream <-chan interface{}, multiplyVal int) <-chan interface{} {
		stream := make(chan interface{})

		go func() {
			defer fmt.Println("exiting multiplier")
			defer close(stream)

			for item := range inStream {
				select {
				case <-done:
					return
				case stream <- item.(int) * multiplyVal:
				}
			}
		}()

		return stream
	}

	input := func() interface{} {
		value := rand.Intn(10) + 1
		fmt.Printf("random: %d\n", value)
		return value
	}

	done := make(chan interface{})

	pipeLine := multiplier(done, generator(done, input), 2)

	count := 1
	for value := range pipeLine {
		fmt.Printf("%d. %d\n", count, value)
		if count == 5 {
			// close(done)
			break
		}
		count++
	}

	close(done)
	time.Sleep(1 * time.Second)
}

func BasicChanThree() {
	randChan := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("rand channel closed")
			defer close(randStream)

			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()

		return randStream
	}

	done := make(chan interface{})
	randStream := randChan(done)

	fmt.Println("3 random values...")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	close(done)

	time.Sleep(1 * time.Second)
}

func BasicChanTwo() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("dowork exited.")
			defer close(completed)

			fmt.Println("before")
			for s := range strings {
				fmt.Println(s)
			}
			fmt.Println("after")
		}()

		return completed
	}

	doWork(nil)
	time.Sleep(2 * time.Second)
	fmt.Println("done")
}

func BasicChanWithOwnership() {
	chOwner := func() <-chan int {
		c := make(chan int, 5)

		go func() {
			defer close(c)

			for i := 1; i <= 10; i++ {
				c <- i
			}
		}()

		return c
	}

	res := chOwner()
	for item := range res {
		fmt.Printf("Receive: %d\n", item)
	}
}

func BasicChan() {
	inStream := make(chan int)

	go func() {
		defer close(inStream)

		for i := 1; i <= 5; i++ {
			inStream <- i
		}
	}()

	for i := range inStream {
		fmt.Println(i)
	}
}

func GoInLoopOne() {
	var wg sync.WaitGroup

	for _, value := range []string{"Zahir", "Ahmed", "Chisty"} {
		wg.Add(1)

		go func(v string) {
			defer wg.Done()

			fmt.Println(v)
		}(value)
	}

	wg.Wait()
}

func DataRace() {
	x := 0

	go func() {
		for {
			x = x + 15
		}
	}()

	go func() {
		for {
			x = x - 15
		}
	}()

	time.Sleep(5 * time.Second)
}
