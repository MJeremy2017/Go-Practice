package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func maxSquare(square []string) int {
	var sqSlice [][]int

	for _, value := range square {
		n := len(value)
		var row []int
		for i:=0; i<n; i++ {
			intV, err := strconv.Atoi(string(value[i]))
			if err!=nil {
				fmt.Println(err)
			}
			row = append(row, intV)
		}
		sqSlice = append(sqSlice, row)
	}
	fmt.Println("Int matrix \n", sqSlice)

	// check square

	var res []int
	for i, row := range sqSlice {
		if i == 0 {
			res = find(1, row)
		}
		res = intersection(find(1, row), res)
	}

	hgt := len(sqSlice)
	lon := longestC(res)

	var t int
	if lon <= hgt {
		t = lon
	} else {t = hgt}

	fmt.Println("res", res)
	return t
}

func find(elem int, slc []int) []int {
	var ind []int
	for i, v := range slc {
		if elem == v {
			ind = append(ind, i)
		}
	}
	return ind
}

func contains(elem int, slc []int) bool {
	for _, v := range slc {
		if elem == v {
			return true
		}
	}
	return false
}

func intersection(slice1 []int, slice2 []int) []int {
	var inter []int
	for _, v := range slice2 {
		if contains(v, slice1) {
			inter = append(inter, v)
		}
	}

	return inter
}

func longestC(slice []int) int {
	// get longest continuous
	var res []int
	var currentV int
	lon := 1
	for i, v := range slice {
		if i == 0 {
			currentV = v
		} else {
			if v - currentV == 1 {
				lon += 1
				currentV = v
			} else {
				res = append(res, lon)
				currentV = v
				lon = 1
			}
		}
	}
	sort.Ints(res)
	fmt.Println("sort res", res)
	return res[len(res)-1]
}

func main() {
	a := []string {"0111", "1101", "0111"}
	fmt.Println(reflect.TypeOf(string(a[0][0])))

	fmt.Println(strconv.Atoi(string(a[0][0])))

	fmt.Println(maxSquare(a))
	fmt.Println("--------------")
	fmt.Println(longestC([]int{1, 2, 5, 6, 7, 9}))
	
}
