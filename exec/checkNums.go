package main

import "fmt"

func lessThan(num1 float32, num2 float32) interface{} {
	if num1 < num2 {
		return true
	} else if num1 > num2 {
		return false
	}
	return -1
}

func main() {
	fmt.Println(lessThan(33, 33))
}
