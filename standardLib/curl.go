package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// create a file to persist response
	file, err := os.Create(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// response write to both file and stdout
	dest := io.MultiWriter(file, os.Stdout)
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
