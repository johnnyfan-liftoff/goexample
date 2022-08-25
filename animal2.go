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
	kind       string
}

var cow, bird, snake Animal = Animal{"cow", "grass", "walk", "moo", "cow"},
	Animal{"bird", "worms", "fly", "peep", "bird"},
	Animal{"snake", "mice", "slither", "hsss", "snake"}

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
		var command string

		fmt.Scanln(&command, &animal, &action)

		switch command {
		case "newanimal":
			fmt.Println("newannimal")
			name := animal
			kind := action
			var ani Animal
			switch kind {
			case "cow":
				ani.name, ani.kind, ani.food, ani.locomotion, ani.noice = cow.name, cow.kind, cow.food, cow.locomotion, cow.noice
			case "bird":
				ani.name, ani.kind, ani.food, ani.locomotion, ani.noice = bird.name, bird.kind, bird.food, bird.locomotion, bird.noice
			case "snake":
				ani.name, ani.kind, ani.food, ani.locomotion, ani.noice = snake.name, snake.kind, snake.food, snake.locomotion, snake.noice
			}

		case "query":
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
		default:
			fmt.Println("Wrong command")
		}

	}
}
