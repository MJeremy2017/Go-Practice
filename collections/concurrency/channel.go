package main

import "fmt"

func main() {
	buffered := make(chan string, 10)
	unbuffered := make(chan string)

	buffered <- "Gopher"
	// receive value from a channel
	value := <-buffered

	fmt.Println(value)
	fmt.Println(unbuffered)
	fmt.Println(buffered)
}