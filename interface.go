package main

import (
	"fmt"
	"math"
)

type gemotry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2*r.height + 2*r.width
}

func (r circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

func (r circle) perim() float64 {
	return 2 * math.Pi * r.radius
}

func measure(g gemotry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{height: 2, width: 3}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
