package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

const maxLength int = 20

func main() {
	var fileName string
	name := make([]Person, 0)
	var person Person

	fmt.Print("Please enter file name:")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		if len(s[0]) > maxLength {
			person.fname = s[0][0:maxLength]
		} else {
			person.fname = s[0]
		}

		if len(s[1]) > maxLength {
			person.lname = s[1][0:maxLength]
		} else {
			person.lname = s[1]
		}
		name = append(name, person)
	}

	for _, v := range name {
		fmt.Println(v.fname, v.lname)
	}
}
