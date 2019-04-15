package main

import (
	"fmt"
	"strings"
)

func longestWord(str string) string {
	lenMax := 0
	stMax := ``
	for _, s:= range strings.Split(str, " ") {
		if len(s) >= lenMax {
			stMax = s
			lenMax = len(s)
		}
	}
	return stMax
}

func longestWord2(str string) string {
	longest := ``
	for _, s:= range strings.Split(str, " ") {
		if len([]rune(longest)) < len([]rune(s)) {
			// rune is used to distinguish character values from integer values
			longest = s
		}
	}
	return longest
}

func main() {
	fmt.Println(len("dsf dfd"))
	fmt.Println(strings.Split("I am J", " "))
	fmt.Println(longestWord("I love dog badly"))
	fmt.Println(longestWord2("I love dog badly"))
}
