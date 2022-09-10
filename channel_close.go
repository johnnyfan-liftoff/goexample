package main

import "fmt"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received", j)
			} else {
				fmt.Println("Done with all the work")
				done <- true
				return
			}
		}
	}()

	// send jobs
	for i := 1; i < 25; i++ {
		jobs <- i
	}
	close(jobs)

	<-done
}
