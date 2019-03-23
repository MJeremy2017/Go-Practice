package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg2 sync.WaitGroup
)

func incCounter(id int) {
	defer wg2.Done()

	for count:=0; count<2; count++ {
		atomic.AddInt64(&counter, 1)
		// let the other goroutine to run
		runtime.Gosched()
	}
}

func main() {
	wg2.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg2.Wait()
	fmt.Println("Final counter", counter)  // 4
}

