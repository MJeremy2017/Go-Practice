package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wgc sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wgc.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1

	wgc.Wait()
	fmt.Println("Game over")
}

func player(name string, court chan int)  {
	defer wgc.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s won\n", name)
			return
		}

		// randomly choose lose condition
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s missed\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++

		court <- ball
	}

}