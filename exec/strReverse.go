package main

import "fmt"

func strReverse(str string) string {
	reverseStr := ""
	for i := len(str)-1; i >=0; i-- {
		reverseStr += string(str[i])
	}
	return reverseStr
}

func main() {
	for _, c := range "I love dogs" {
		fmt.Println(string(c))
	}

	fmt.Println("c" + "a")
	fmt.Println( string("dadf"[3]) )

	fmt.Println(strReverse("I love dogs"))
}