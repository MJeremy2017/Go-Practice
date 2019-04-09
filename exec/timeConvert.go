package main

import (
	"fmt"
	"strconv"
)

func convertTime(num int) string {
	hour := num/60
	min := num%60
	r := strconv.Itoa(hour) + ":" + strconv.Itoa(min)
	return r
}

func main()  {
	fmt.Println(convertTime(126))
	fmt.Println(string(2))
}