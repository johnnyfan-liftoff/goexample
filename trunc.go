package main

import "fmt"

func main() {
	fmt.Println("Please input a floating point number:")
	var number float32
	fmt.Scan(&number)
	fmt.Println("The truncated number is:", int32(number))
}
