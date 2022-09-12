package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println("Hello")
	fmt.Println(m)

	m["c"] = 3
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
