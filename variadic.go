package main

import "fmt"

func sum(num ...int) int {

	var total int
	for _, i := range num {
		total += i
	}
	return total
}

func main() {

	num := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(num...))
}
