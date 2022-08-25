package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	id int
}

type Philosopher struct {
	id int
}

var Pool = sync.Pool{
	New: func() interface{} {
		return new(Chopstick)
	},
}

func (a Philosopher) eat(hosts chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	<-hosts

	fmt.Printf("starting to eat %d\n", a.id)
	left := Pool.Get()
	right := Pool.Get()

	fmt.Printf("finishing eating %d\n", a.id)
	Pool.Put(left)
	Pool.Put(right)

	hosts <- 1
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		Pool.Put(new(Chopstick))
	}

	hosts := make(chan int, 2)

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			wg.Add(1)
			go philosophers[i].eat(hosts, &wg)
		}
	}

	hosts <- 1
	hosts <- 1

	wg.Wait()
}
