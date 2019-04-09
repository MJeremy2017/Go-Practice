package main

import "fmt"

func symbols(str string) bool {

	if str[0] >= 'a' && str[0] <= 'z' {
		return false
	}

	if str[len(str)-1] >= 'a' && str[len(str)-1] <= 'z' {
		return false
	}
	
	flag := true
	for i := 1; i < len(str)-1; i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			if str[i-1] != '+' || str[i+1] != '+' {
				return false
			}
		}
	}
	return flag
}

func main() {
	fmt.Println(symbols("+d+=3=+s+"))
	fmt.Println(symbols("f++d+"))
}