package main

import "fmt"

func main() {

	const s = "สวัสดี"
	// const s = "jjjjjbbbb"
	fmt.Println("len=", len(s))

	for i, v := range s {
		fmt.Printf("++%d++: --%#U--", i, v)
	}
	fmt.Println("")

	for j := 1; j < len(s); j++ {
		fmt.Printf("%x", s[j])
	}

}
