package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter2 int
	wtg sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wtg.Add(2)

	go incCounter2(1)
	go incCounter2(2)

	wtg.Wait()
	fmt.Println("Final counter", counter2)
}

func incCounter2(id int) {
	defer wtg.Done()

	for count:=0; count<2; count++ {
		mutex.Lock()  // lock the following code

		value := counter2
		runtime.Gosched()
		value ++

		counter2 = value

		mutex.Unlock() // each time only 1 goroutine can access the code
	}
}