package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		fmt.Println("inside func")
		i++
		return i
	}
}

func main() {
	a := intSeq()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}
