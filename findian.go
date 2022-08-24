package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	println("Please enter a string:")
	fmt.Scan(&str)
	str = strings.ToLower(str)

	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
