package main

import (
	"errors"
	"fmt"
)

func f1(num int) (int, error) {
	if num == 0 {
		return -1, errors.New("Error : 0")
	}
	return num + 1, nil
}

func f2(num int) (int, error) {
	if num == 0 {
		return -1, &argError{arg: 1, prob: "zero error issue"}
	}
	return num + 1, nil
}

type argError struct {
	arg  int
	prob string
}

func (arg *argError) Error() string {
	return fmt.Sprintf("Error number: %d Prob: %s", arg.arg, arg.prob)
}

func main() {
	num, err := f1(22)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	num, err = f1(0)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	num, err = f2(22)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	num, err = f2(0)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
