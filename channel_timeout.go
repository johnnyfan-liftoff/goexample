package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	ch1 := make(chan string, 5)
	done := make(chan bool)

	go func() {
		fmt.Println("world")

		select {
		case msg := <-ch1:
			fmt.Println("aa", msg)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")
			// default:
			// 	fmt.Println("Didn't get anything")
		}
		fmt.Println("Done func")
		done <- true
	}()

	// ch1 <- "message"

	<-done
	fmt.Println("Done")
}
