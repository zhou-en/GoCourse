package main

import (
	"fmt"
	// "math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base float64
	height float64
}

type square struct {
	side float64
}

func main() {

	var tri triangle
	tri.height = 2.0
	tri.base = 1.0

	var squ square
	squ.side = 2.5

	printArea(tri)
	printArea(squ)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}
