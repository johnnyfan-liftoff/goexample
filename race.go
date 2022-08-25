package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup

// This routine is to count 1000000 times and save it into global variable x
func count() {
	defer wg.Done()
	fmt.Println("start")
	i := 0
	for i < 1000000 {
		x++
		i++
	}
	fmt.Println("finish")
}

// The routine count() was called twice to count for 1000000 times. Since this routine
// was called twice, the total count should be 2000000. However, due to the race condition,
// when the first routine is counting, the second routine gets executed at the same time.
// The second routine may overwrite the value of x to another number, which cause the total
// number is less than 2000000.

// Execution output:
// C:\goexample>go run race.go
// start
// start
// finish
// finish
// Total count: 1702847

func main() {

	wg.Add(2)
	go count()
	go count()

	wg.Wait()
	fmt.Println("Total count:", x)
}
