package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base float64
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Printf("Area is %v\n", s.getArea() )
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}



func main() {
	t := triangle {
		base: 10.0,
		height: 20.0,
	}
	printArea(t)

	s := square{
		sideLength: 5.0,
	}
	printArea(s)


}