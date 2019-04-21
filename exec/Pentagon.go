package main

import "fmt"

func pentagonNum(x int) int {
	if x == 1 {
		return 1
	}

	return x*5-5 + pentagonNum(x-1)
}

func main()  {
	fmt.Println(pentagonNum(2))
	fmt.Println(pentagonNum(3))
	fmt.Println(pentagonNum(5))
}
