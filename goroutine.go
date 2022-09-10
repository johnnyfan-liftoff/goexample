package main

import (
	"fmt"
	"time"
)

func f(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, i)
	}
}

func main() {
	f("direct")

	go f("routine")

	go func(name string) {
		fmt.Println(name)
	}("anonymous")

	time.Sleep(time.Second)
	fmt.Println("Done")
}
