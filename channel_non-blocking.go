package main

import "fmt"

func main() {

	ch1 := make(chan string, 1)
	ch2 := make(chan string)
	sig := make(chan bool)

	go func() {

		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		default:
			fmt.Println("nothing")
		}
		fmt.Println("Done")
		sig <- true
		return
	}()

	ch2 <- "world"
	fmt.Println("Done in main - world")
	ch1 <- "hello"
	fmt.Println("Done in main - hello")
	close(ch1)
	fmt.Println("Done in main - close1")
	close(ch2)
	fmt.Println("Done in main - close2")

	<-sig
	fmt.Println("Done in main")
}
