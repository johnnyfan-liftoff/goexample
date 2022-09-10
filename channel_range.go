package main

import "fmt"

func main() {

	ch1 := make(chan string, 5)

	go func() {
		ch1 <- "Hello1"
		ch1 <- "Hello2"
		ch1 <- "Hello3"
		ch1 <- "Hello4"
		ch1 <- "Hello5"
		close(ch1)
	}()

	for i := range ch1 {
		fmt.Println(i)
	}
}
