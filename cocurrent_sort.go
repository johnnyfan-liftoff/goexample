package main

import (
	"fmt"
	"sort"
	"sync"
)

func chunks(slice []int, size int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func main() {
	var a []int

	fmt.Print("Please input numbers to sort:")
	for {
		var t int
		fmt.Scan(&t)
		if t == -1 {
			break
		}
		a = append(a, t)
	}

	size := len(a)/4 + 1
	parts := chunks(a, size)

	var w sync.WaitGroup

	for i := 0; i < 4; i++ {

		w.Add(1)

		go func(arr []int) {
			fmt.Println("part: ", arr)
			sort.Ints(arr)

			w.Done()
		}(parts[i])
	}

	w.Wait()

	sort.Ints(a)
	fmt.Println("sorted: ", a)
}
