package main

import "fmt"

func isValid(s string) bool {
	buf := make([]byte, 0)
	for _, c := range s {
		switch c {
		case '(':
			buf = append(buf, ')')
		case '[':
			buf = append(buf, ']')
		case '{':
			buf = append(buf, '}')
		case ')', ']', '}':
			if len(buf) == 0 {
				return false
			}
			if rune(buf[len(buf)-1]) == rune(c) {
				buf = buf[:len(buf)-1]
			} else {
				return false
			}
		}
	}

	if len(buf) == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	s := "()"
	fmt.Println(isValid(s))
}
