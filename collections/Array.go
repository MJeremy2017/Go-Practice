package main

import "fmt"

func main() {
	var array [5]int  // initialise an array of zeros
	array2 := [...]int{2, 1, 2, 3, 5}  // array literals
	array3 := [5]int{1: 22, 2: 3}

	// reassigning
	array3[3] = 11
	// array of pointers
	array4 := [5]*int{0: new(int), 1: new(int)}
	*array4[0] = 10
	fmt.Println(*array4[0])  // value & is address

	// copy array
	arrString1 := [3]string{"a", "b", "c"}
	var arrString2 [3]string
	arrString2 = arrString1

	// copy array of pointers
	var arrPters1 [3]*string
	arrPters2 := [3]*string{new(string), new(string), new(string)}
	*arrPters2[0] = "a"
	*arrPters2[1] = "b"
	arrPters1 = arrPters2

	// multidimensional array
	arrMul := [3][2]int{{1, 1}, {2, 2}, {3, 3}}
	arrMul[2][0] = 10
	var value = arrMul[1][1]

	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(arrString2)
	fmt.Println(arrPters1)
	fmt.Println(arrMul)
	fmt.Println(arrMul[1])  // second row of matrix
	fmt.Println(value)
}
