package main

import "fmt"

func func1(i int) func() int {

	fmt.Println(i)
	return func() int {
		fmt.Println("inner func")
		return 2
	}
}

func main() {
	a := func1(2)
	a()
}
