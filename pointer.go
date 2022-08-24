package main

import "fmt"

func main() {
	var x int = 1
	var y int
	var ip *int

	ip = &x
	y = *ip

	fmt.Println(ip)
	fmt.Println(&x)
	fmt.Println(y)

	ptr := new(int)
	*ptr = 3
	fmt.Println(*ptr)
}
