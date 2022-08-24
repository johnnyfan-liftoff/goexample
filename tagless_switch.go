package main

import "fmt"

func main() {
	var x int
	// var appleNum int
	fmt.Printf("Number of apples?")
	n, err := fmt.Scan(&x)

	fmt.Println(n, err)

	switch {
	case x > 1:
		fmt.Println("case1")
	case x < -1:
		fmt.Println("case2")
	default:
		fmt.Println("nocase")
	}
}
