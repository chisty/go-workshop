package main

import (
	"fmt"
	"runtime"
	"sync"
)

func GoMemConsumedOne() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup

	noop := func() {
		wg.Done()
		<-c
	}

	const numGoroutines = 10000
	wg.Add(numGoroutines)

	before := memConsumed()

	for i := numGoroutines; i > 0; i-- {
		go noop()
	}

	wg.Wait()

	after := memConsumed()

	fmt.Printf("Total memory: %.3fmb\n", float64(after-before)/100000)

	fmt.Printf("Each goroutine memory(avg): %.3fkb", float64(after-before)/numGoroutines/1000)
}
