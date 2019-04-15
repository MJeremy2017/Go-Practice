package main

import (
	"fmt"
	"strings"
)

func letChange(str string) string {
	vowels := "aeiou"
	alphas := "abcdefghijklmnopqrstuvwxyz"
	res := ``
	for _, l := range str {
		nxt := ""
		newChar := strings.ToLower(string(l))
		if strings.Index(alphas, newChar) < 0 {
			nxt = newChar
		} else if newChar == `z` {
			nxt = `a`
		} else if strings.Index(vowels, newChar) >= 0 {
			nxt = strings.ToUpper(string(int(l) + 1))
		} else {
			nxt = string(int(l) + 1)
		}
		res += nxt
	}

	return res
}

func main() {
	s := letChange("I love dogs z!")
	fmt.Println(s)
	fmt.Println(strings.Index("adsadf", "d"))
}
