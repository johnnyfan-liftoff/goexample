package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

var cow, bird, snake Animal = Animal{"cow", "grass", "walk", "moo"},
	Animal{"bird", "worms", "fly", "peep"},
	Animal{"snake", "mice", "slither", "hsss"}

func (animal *Animal) Eat() {
	fmt.Printf("The %s eats %s\n", animal.name, animal.food)
}

func (animal *Animal) Move() {
	fmt.Printf("The %s %s\n", animal.name, animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Printf("The %s %s\n", animal.name, animal.noise)
}

func (animal *Animal) Perform(action string) {
	switch strings.ToLower(action) {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Wrong action")
	}
}

func main() {
	for {
		fmt.Print(">")
		var animal string
		var action string

		fmt.Scanln(&animal, &action)

		switch animal {
		case "cow":
			cow.Perform(action)
		case "bird":
			bird.Perform(action)
		case "snake":
			snake.Perform(action)
		default:
			fmt.Println("Wrong animal")
		}
	}
}
