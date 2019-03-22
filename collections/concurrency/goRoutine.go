package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {
	// max number of logical processors
	runtime.GOMAXPROCS(2) // change to 2

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start goroutine")

	// create the first goroutine
	go func() {
		// tell main to wait
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println()
		}
		// wg.Done()

	}()

	// create the second goroutine
	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println("")
		}
		//wg.Done()
	}()

	// wait the goroutine done
	fmt.Println("Waiting ...")
	wg.Wait()

	fmt.Println("Done!")
}
