package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortAlpha(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	x := "klda"
	s := strings.Split(x, "")
	fmt.Println(s)

	fmt.Println(sortAlpha(x))
}

