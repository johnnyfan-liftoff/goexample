package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Array
	array := [3]int{1, 2, 3}
	arrayCopy := array
	fmt.Println("type:", reflect.TypeOf(array))
	fmt.Printf("&array[0]: %p\n", &array[0])
	fmt.Printf("&arrayCopy[0]: %p\n", &arrayCopy[0])

	fmt.Println("---------------------------")

	// Slice
	slice := []int{1, 2, 3}
	sliceRef := slice
	fmt.Println("type:", reflect.TypeOf(slice))
	fmt.Printf("&slice[0]: %p\n", &slice[0])
	fmt.Printf("&sliceRef[0]: %p\n", &sliceRef[0])
	slice = append(slice, 4, 5)
	fmt.Printf("&slice[0]: %p\n", &slice[0])
	fmt.Printf("&sliceRef[0]: %p\n", &sliceRef[0])
}
