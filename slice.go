package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var list []int = make([]int, 3)
	var str string

	for {
		fmt.Println("Input an interger(input 'X' to exit):")
		fmt.Scanln(&str)
		if str == "X" {
			break
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Wrong format")
			continue
		}
		list = append(list, num)
		sort.Ints(list[:])
		fmt.Println(list)
	}
}
