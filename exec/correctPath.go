package main

import (
	"fmt"
	"strings"
)

func position(x int, y int, action string) (int, int) {

	switch action {
	case "l":
		return x, y-1
	case "r":
		return x, y+1
	case "u":
		return x-1, y
	case "d":
		return x+1, y
	}
	return -1, -1
}

func getPath(p string) {
	numQ := strings.Count(p, "?")
	numL := strings.Count(p, "l")
	numR := strings.Count(p, "r")
	numU := strings.Count(p, "u")
	numD := strings.Count(p, "d")

	numD = numD - numU
	numR = numR - numL

	neededD := 4 - numD
	neededR := 4 - numR

	var replaceString []string
	if neededD>0 {
		for i:=0;i<neededD;i++ {
			replaceString = append(replaceString, "d")
		}
	} else {
		for i:=0;i<(-neededD);i++ {
			replaceString = append(replaceString, "u")
		}
	}

	if neededR>0 {
		for i:=0;i<neededR;i++ {
			replaceString = append(replaceString, "r")
		}
	} else {
		for i:=0;i<(-neededR);i++ {
			replaceString = append(replaceString, "l")
		}
	}

	for i:=0; i<numQ-len(replaceString)+1; i++ {
		if i%2==0 {
			replaceString = append(replaceString, "u")
		} else {
			replaceString = append(replaceString, "d")
		}
	}

	fmt.Println("replacement", replaceString)
	var res []string
	i := 0
	for _, r := range p {
		if string(r)=="?" {
			res = append(res, replaceString[i])
			i += 1
		} else {
			res = append(res, string(r))
		}
	}

	fmt.Println(res)

}

func main() {
	fmt.Println(position(3, 2, "r"))
	fmt.Println(strings.Count("asssss", "s"))
	getPath("???rrurdr?")
	getPath("drdr??rrddd?")
}
