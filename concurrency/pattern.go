package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func WorkerPool() {
	numWorker := 10
	numJobs := 100
	resMap := make(map[int][]int)
	var resMapMutex sync.Mutex

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	type resData struct {
		worker int
		job    int
	}

	jobs := make(chan int)
	result := make(chan resData, numJobs)

	var wg sync.WaitGroup
	var resWg sync.WaitGroup

	for i := 1; i <= numWorker; i++ {
		wg.Add(1)

		go func(id int, jobs <-chan int) {
			defer wg.Done()
			for job := range jobs {
				select {
				case <-ctx.Done():
					return
				default:
					time.Sleep(3 * time.Millisecond)
					result <- resData{worker: id, job: job}
				}
			}

		}(i, jobs)
	}

	resWg.Add(1)
	go func() {
		defer resWg.Done()

		for res := range result {
			resMapMutex.Lock()

			if _, ok := resMap[res.worker]; !ok {
				resMap[res.worker] = []int{}
			}
			resMap[res.worker] = append(resMap[res.worker], res.job)

			resMapMutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= numJobs; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
			}
		}

		close(jobs)
		wg.Wait()
		close(result)
	}()

	resWg.Wait() //wait for the result goroutine

	for key, val := range resMap {
		fmt.Printf("Worker: %d, Count: %d, Jobs: %v\n", key, len(val), val)
	}

	fmt.Println("Process complete")
}

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()

		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))

	for _, c := range channels {
		go multiplex(c)
	}

	return multiplexedStream
}
