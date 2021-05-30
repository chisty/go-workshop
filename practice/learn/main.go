package main

import (
	"fmt"
	"math/rand"
)

func main() {
	doneChan := make(chan bool)
	stream := getProducer(doneChan)

	for i := 0; i < 5; i++ {
		fmt.Printf("%d. Value %d\n", i+1, <-stream)
	}

	close(doneChan)

	fmt.Println("Done!")
}

func getProducer(done <-chan bool) <-chan int {
	stream := make(chan int)

	go func() {
		defer fmt.Println("Producer stopped!")
		defer close(stream)

		for {
			select {
			case stream <- rand.Int():
			case <-done:
				return
			}
		}
	}()
	return stream
}
