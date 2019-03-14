package main

import "fmt"

func main()  {
	slice := make([]string, 5)
	slice2 := make([]int, 3, 5)  // length 3 capacity 5
	slice3 := []int{1, 2, 3}
	var slice4 []int  // or slice := []int{}

	// slicing
	orgSlice := []int{1, 2, 3, 4, 5}
	newSlice := orgSlice[1:3]  // length = 3-1 = 2; capacity = 5-1 = 4
	newSlice[1] = 333
	newSlice = append(newSlice, 6666)

	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	sliceSource := source[2:3:5]
	// restrict capacity
	sliceSource = source[2:3:3]
	sliceSource = append(sliceSource, "Kiwi")

	slice6 := []int{1, 2}
	slice6 = append(slice6, slice6...)

	newSlice = []int{1, 2, 3}
	for index, value := range newSlice {
		fmt.Printf("index: %d  value: %d \n", index, value)
	}

	for i := 0; i < len(newSlice); i++  {
		fmt.Printf("index: %d value: %d\n", i, newSlice[i])
	}

	// multi-dimensional slice
	mulSlice := [][]int{{10}, {100, 200}}
	fmt.Println(mulSlice)
	mulSlice[0] = append(mulSlice[0], 20)
	fmt.Println(mulSlice)


	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(newSlice)
	fmt.Println(orgSlice)
	fmt.Println("after appending ...")
	fmt.Println(newSlice)
	fmt.Println(orgSlice)
	fmt.Println(sliceSource)
	fmt.Println(slice6)
}
