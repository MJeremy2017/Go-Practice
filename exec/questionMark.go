package main

import (
	"fmt"
	"strconv"
)

func question(st string) {
	qCount := 0
	startCount := false
	currentSum := 0
	for _, i := range st {
		if string(i) == "?" {
			if startCount {
				qCount += 1
			}
		}

		if (i >= '0') && (i <= '9') {
			t, err := strconv.Atoi(string(i))
			if err != nil {
				fmt.Println(err)
			}
			currentSum += t
			startCount = !startCount

			if !startCount && currentSum == 10 {
				if qCount == 3 {
					fmt.Println("true")
					return
				} else {
					currentSum = 0
				}
			}
			qCount = 0
		}
	}
	fmt.Println("false")
}

func main() {
	question("s??2???3")
	fmt.Println()
	question("acc?7??sss?3rr1??????5")
}
