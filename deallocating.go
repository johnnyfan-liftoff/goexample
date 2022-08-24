package main

import "fmt"

var x = 4 // in heap

func main() {
	x := 0 // in stack
	fmt.Println(x)

	i := 0
	// x is created for 10 times.
	for i < 10 {
		f()
		i++
	}

	var k float32 = 123.45           // 6 digits of precision
	var l float64 = 1.2345e2         // 15 digits of precision
	var z complex128 = complex(2, 3) // complex numbers
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(z)
}

func f() {
	var x = 4
	fmt.Printf("%d", x)
	fmt.Println("")
}
