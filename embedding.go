package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num %d", b.num)
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "message",
	}

	fmt.Println("co:", co.num, co.str)
	fmt.Println("co.base", co.base.num)
	fmt.Println("co.describe", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("d describe()", d.describe())

}
