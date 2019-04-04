package main

import (
	"github.com/goinaction/code/chapter9/listing17/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Routes()

	log.Println("listener: Listening on port : 4000")
	http.ListenAndServe(":4000", nil)
}
