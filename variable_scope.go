package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	f()
	g()
	h()
	i()
}

var x = 4

func f() {
	fmt.Printf("%d", x)
}

func g() {
	fmt.Printf("%d", x)
}

func h() {
	var x = 4
	fmt.Printf("%d", x)
}

func i() {
	fmt.Printf("%d", x)
}
