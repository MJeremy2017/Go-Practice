package main

import "fmt"

func chessWalk(str string) int {
	var nums []int

	for _, s := range str {
		if s >= '0' && s <= '9' {
			nums = append(nums, int(s)-'0')
		}
	}

	h := nums[2] - nums[0]
	w := nums[3] - nums[1]

	return factorial2(h+w)/(factorial2(h)*factorial2(w))
}

func factorial2(num int) int {
	if num==1 {
		return 1
	}

	return num*factorial2(num-1)
}

func main() {
	fmt.Println(chessWalk("(1, 1)(4, 4)"))
}
