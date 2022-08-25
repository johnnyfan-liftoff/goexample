package main

import "fmt"

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) (s float64) {
		s = .5*a*t*t + v0*t + s0
		return
	}
	return fn
}

func main() {

	var a, v0, s0, t float64

	fmt.Print("Please input acceleration:")
	fmt.Scan(&a)

	fmt.Print("Please input initial velocity:")
	fmt.Scan(&v0)

	fmt.Print("Please input initial displacement:")
	fmt.Scan(&s0)

	fmt.Print("Please input time interval:")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Printf("Displacement after %.2f seconds: %.4f.\n", t, fn(t))
}
