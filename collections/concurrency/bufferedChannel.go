package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad = 10
)

var wgg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// buffered
	tasks := make(chan string, taskLoad)

	wgg.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// add work to be done
	for post := 1; post <= 10; post++ {
		tasks <- fmt.Sprintf("Task -> %d", post)
	}

	close(tasks)

	wgg.Wait()

}

func worker(tasks chan string, worker int)  {
	defer wgg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			// means channel is empty
			fmt.Printf("Worker %d shutting down \n", worker)
			return
		}

		fmt.Printf("Worker: %d working Task: %s \n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d completed Task: %s \n", worker, task)
	}
}

