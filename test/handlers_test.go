package main

import (
	"encoding/json"
	"github.com/goinaction/code/chapter9/listing17/handlers"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const ballotX2 = "\u2717"
const checkMark2 = "\u2713"

func init()  {
	handlers.Routes()
}

func TestSendJson(t *testing.T) {
	t.Log("Given the need to test sendJson")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err!=nil {
			t.Fatal("Request error", ballotX2)
		}
		t.Log("Able to get request", checkMark2)

		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)  // send to server and get response

		if rw.Code!=200 {
			log.Fatal("Error response", ballotX2, rw.Code)
		}
		t.Log("Get response", checkMark2)


		u := struct {
			Name string
			Email string
		}{}
		t.Log("\tShould decode the response.", checkMark2)

		if err:=json.NewDecoder(rw.Body).Decode(&u); err!=nil {
			t.Fatal("Error decode response", ballotX2, err)
		}

		t.Log("Able to decode response", checkMark2)

		if u.Name == "Bill" {
			t.Log("\tShould have a Name.", checkMark2)
		} else {
			t.Error("\tShould have a Name.", ballotX2, u.Name)
		}


		if u.Email == "bill@ardanstudios.com" {
			t.Log("\tShould have an Email.", checkMark2)
		} else {
			t.Error("\tShould have an Email.", ballotX2, u.Email)
		}

	}
}