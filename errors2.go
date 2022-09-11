package main

import (
	"errors"
	"fmt"
)

type argError struct {
	err int
	msg string
}

func (e *argError) Error() string {
	return "arg error"
}

func f2(b int) (int, error) {
	fmt.Println("f2 error")
	arg := argError{33, "dddd"}
	return 88, &arg
}

func f1(a int) (int, error) {
	fmt.Println("f1")
	return 55, errors.New("new error message")
}
func main() {
	a, e := f1(33)
	fmt.Println(a, e)

	b, e2 := f2(99)
	fmt.Println(b, e2)
}
