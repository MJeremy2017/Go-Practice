package main

import (
	"fmt"
	"strings"
)

func capitalize(str string) string {
	alphas := "abcdefghijklmnopqrstuvwxyz"
	isFirst := true
	res := ``
	for _, l := range str {
		if strings.Index(alphas, string(l))>0 {
			if isFirst {
				res += strings.ToUpper(string(l))
			} else {
				res += string(l)
			}
			isFirst = false
		} else {
			res += string(l)
			isFirst = true
		}
	}
	return res
}

func main() {
	s := capitalize("hello# world")
	fmt.Println(s)
}
