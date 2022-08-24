package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]string)

	in := bufio.NewScanner(os.Stdin)

	fmt.Print("Please input your name:")
	in.Scan()
	m["name"] = in.Text()

	fmt.Print("Please input your address: ")
	in.Scan()
	m["address"] = in.Text()

	obj, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(obj))
}
