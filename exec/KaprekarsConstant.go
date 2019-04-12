package main

import (
	"fmt"
	"sort"
)

func kaprekars(num int) int {

	if num == 6174 {
		return 0
	}

	nl := []int{num / 1000, num / 100 % 10, num / 10 % 10, num % 10}

	sort.Ints(nl)

	ord := nl[0]*1000 + nl[1]*100 + nl[2]*10 + nl[3]
	rev := nl[3]*1000 + nl[2]*100 + nl[1]*10 + nl[0]

	return kaprekars(rev-ord) + 1
}

func main() {
	nl := []int{3, 4, 1}
	sort.Ints(nl)
	fmt.Println(nl)

	fmt.Println(4321/100 % 10)

	fmt.Println(kaprekars(2111))
}
