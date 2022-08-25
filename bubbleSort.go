package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(a []int, i int) {
	a[i], a[i+1] = a[i+1], a[i]
}

func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j+1] < a[j] {
				Swap(a, j)
			}
		}
	}
}

func Input() (values []int, err error) {
	fmt.Println("Please input:")
	bf := bufio.NewReader(os.Stdin)
	line, _, err := bf.ReadLine()

	tokens := strings.Split(string(line), " ")
	for _, s := range tokens {
		n, _ := strconv.Atoi(s)
		values = append(values, n)
	}
	return values, err
}

func main() {
	values, _ := Input()
	BubbleSort(values)
	fmt.Println("Sorted:", values)
}
