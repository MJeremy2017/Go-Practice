package main

import "fmt"

func checkWin(slice [][]int)  bool{
	// check row and col
	var rowInd []int
	var colInd []int
	for _, row := range slice {
		if contain(rowInd, row[0]) {
			return false
		}
		
		rowInd = append(rowInd, row[0])

		if !contain(colInd, row[1]) {
			return false
		}
		
		colInd = append(colInd, row[1])
	}
	
	// check diagonal: whenever i, j increase by 1
	for _, i := range slice {
		for _, j := range slice {
			if (i[0]+1 == j[0]) && (i[1]+1 == j[1]) {
				return false
			}
		}
	}

	return true
	

}

func contain(slice []int, i int) bool{
	for _, v := range slice {
		if i == v {
			return true
		}
	}
	return false
}

func main() {
	s := [][]int{{1, 2}, {3, 4}}
	fmt.Println(s)
	fmt.Println(s[0][0])
	fmt.Println(!true)
	fmt.Println(checkWin([][]int{{2,1}, {4,3}, {6,3}, {8,4}, {3,4}, {1,6}, {7,7}, {5,8}}))
}
