package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Triangle struct {
	width float64
	height float64
}

type Square struct {
	width float64
	height float64
}

func (sq Square) getArea() float64 {
	return sq.width * sq.height
}

func (tr Triangle) getArea() float64 {
	return tr.width * tr.height/2
}

func printArea(shape Shape) {
	fmt.Println("Area is ", shape.getArea())
}

func main() {
	triangle := Triangle{width: 3.2, height: 4.1}
	square := Square{width: 3.2, height: 4.1}

	printArea(triangle)
	printArea(square)

}
