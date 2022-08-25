package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

type AnimalInterface interface {
	Eat()
	Move()
	Speak()
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var animal AnimalInterface
	var command, which, kind string

	aMap := make(map[string]Animal)

	aMap["cow"] = Animal{"grass", "walk", "moo"}
	aMap["bird"] = Animal{"worms", "fly", "peep"}
	aMap["snake"] = Animal{"mice", "slither", "hsss"}

	for {
		fmt.Print(">")
		fmt.Scan(&command, &which, &kind)
		switch command {

		case "query":
			animal = aMap[which]
			switch kind {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		case "newanimal":
			aMap[which] = aMap[kind]
			fmt.Println("Created it!")
		}
	}
}
