package main

import (
	"fmt"
)

func main() {
	dict := make(map[string]string)
	fmt.Println(dict)

	dict = map[string]string{"a": "b", "b": "c"}
	fmt.Println(dict)

	colors := map[string]string{}  // has to name {}
	colors["red"] = "aasfsf"
	fmt.Println(colors)

	// test key existence
	value, exists := colors["blue"]
	fmt.Println(value, exists)  // value = ""

	colors = map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// remove a key
	delete(colors, "Coral")

	for key, value := range colors{
		fmt.Printf("key: %s  value: %s \n", key, value)
	}

	removeColor(colors, "DarkGray")
	for key, value := range colors{
		fmt.Printf("key: %s  value: %s \n", key, value)
	}

}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)  // changes directly applies on original map
}
