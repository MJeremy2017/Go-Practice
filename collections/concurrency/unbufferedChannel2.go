package main

import (
	"fmt"
	"sync"
	"time"
)

var wgb sync.WaitGroup

func main()  {
	baton := make(chan int)

	wgb.Add(1)

	go Runner(baton)
	
	baton <- 1
	
	wgb.Wait()
}

func Runner(baton chan int)  {
	var newRunner int
	runner := <-baton

	fmt.Printf("Runner %d with baton\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the Line\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(100*time.Millisecond)

	if runner == 4 {
		fmt.Println("Race finished")
		wgb.Done()  // last one to signify done
		return
	}

	// exchange baton
	fmt.Printf("Runner %d exchange with runner %d\n", runner, newRunner)
	baton <- newRunner
}