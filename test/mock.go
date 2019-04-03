package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

var feed = `<...>`  // some content

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func main() {
	server := mockServer()
	fmt.Println(server.URL)
}


