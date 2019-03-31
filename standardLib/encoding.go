package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	c := make(map[string]interface{})
	fmt.Println(c)

	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": 1234,
		"cell": 4321,
	}

	// marshal map into json
	data, err := json.MarshalIndent(c, "", "	")

	if err!=nil {
		log.Println(err)
	}

	fmt.Println(string(data))
}