package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		ch <- "ping"
	}()

	fmt.Println("received", <-ch)

	// buffered channel
	buf := make(chan string, 2)
	go func() {
		buf <- "message1"
		time.Sleep(2 * time.Second)
		buf <- "message2"
	}()

	fmt.Println("received", <-buf)
	fmt.Println("received", <-buf)

	// no-blocking channels
	messages := make(chan string)
	singals := make(chan bool)

	go func() {

		messages <- "Hello"
		singals <- true

		<-messages

	}()

	time.Sleep(1 * time.Microsecond)

	select {
	case msg := <-messages:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received", msg)
	case sig := <-singals:
		fmt.Println("Received", sig)
	default:
		fmt.Println("Nothing received")
	}

}
