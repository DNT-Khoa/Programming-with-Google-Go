package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	// Hard code each animal with their properties
	dictionary := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		words := strings.Fields(scanner.Text())

		if len(words) != 2 {
			fmt.Println("Please enter two and only two strings in your request")
			continue
		}

		name := words[0]
		info := words[1]

		details := dictionary[name]
		if details == (Animal{}) {
			fmt.Println("Unknown name")
		}

		switch info {
		case "eat":
			details.Eat()
		case "move":
			details.Move()
		case "speak":
			details.Speak()
		default:
			fmt.Println("Unknown information")
		}
	}
}
