package main

import "fmt"

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	ch1 <- "hello"
	ch2 <- "world"

	a1 := <-ch1
	a2 := <-ch2

	fmt.Println(a1, a2)
}
