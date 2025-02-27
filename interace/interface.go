package main

import (
	"fmt"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {

	var r shape = rectangle{width: 5, height: 3}
	fmt.Println("Rectangle width:", r.area())
	fmt.Println("Rectangle perimeter:", r.perimeter())
}
