package main

import "fmt"

func balance(left []int, right []int) (int, int){
	for _, v1 := range right {
		for _, v2 := range right{
			if left[0]+v1 == left[1]+v2 {
				return v1, v2
			}
		}
	}
	return 0, 0
}

// need imp
func combine(slice []int, n int) []int {
	if n == 1 {
		for _, v := range slice {
			return []int{v}
		}
	}
	for i, v := range slice {
		newSlice := append(slice[:i], slice[i+1:]...)
		return append([]int{v}, combine(newSlice, n-1)...)
	}

	return []int{}
}



func main() {
	fmt.Println(balance([]int{3, 4}, []int{1, 2, 7, 7}))
	fmt.Println(append([]int{3}, []int{3, 4}...))
	fmt.Println(combine([]int{1, 2, 3}, 2))
}