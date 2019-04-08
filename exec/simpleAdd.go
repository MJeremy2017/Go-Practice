package main

import "fmt"

func addSum(x int) int {
	if x == 0 {
		return 0
	} else {
		return x + addSum(x-1)
	}
}

func addSum2(x int) int {
	s := 0
	for i := 0; i <= x; i++ {
		s += i
	}
	return s
}

func main() {
	s := addSum(10)
	fmt.Println(s)

	s2 := addSum2(10)
	fmt.Println(s2)
}