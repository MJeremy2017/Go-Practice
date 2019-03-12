package main

import (
	"fmt"
	"github.com/goinaction/code/chapter3/words"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello World")  // godoc fmt -- more info

	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err!=nil {
		fmt.Println(err)
		return
	}

	text := string(contents)
	count := words.CountWords(text)

	fmt.Printf("There are %d words in the txt \n", count)

}