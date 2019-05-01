package main

import (
	"fmt"
	"strings"
)

func strUpper(strSlice []string) []string {
	var res []string
	for _, s := range strSlice {
		res = append(res, strings.ToUpper(s))
	}

	return res
}

func main() {
	fmt.Println(strings.ToUpper("sda"))
	fmt.Println(strUpper([]string{"as", "Seae3"}))
}
