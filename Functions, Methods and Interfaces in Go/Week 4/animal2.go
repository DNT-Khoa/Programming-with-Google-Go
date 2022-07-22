package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := map[string]Animal{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		words := strings.Fields(scanner.Text())

		if len(words) != 3 {
			fmt.Println("Please enter 3 strings separated by spaces!")
			continue
		}

		action := words[0]
		name := words[1]
		if action == "newanimal" {
			category := words[2]

			switch category {
			case "cow":
				animals[name] = Cow{}
			case "bird":
				animals[name] = Bird{}
			case "snake":
				animals[name] = Snake{}
			default:
				fmt.Println("Unknown animal type!")
				continue
			}

			fmt.Println("Created it!")
		} else if action == "query" {
			locomotion := words[2]
			selectedAnimal := animals[name]
			if selectedAnimal == nil {
				fmt.Printf("Cannot find the animal with name: %s\n", name)
				continue
			}

			switch locomotion {
			case "eat":
				selectedAnimal.Eat()
			case "move":
				selectedAnimal.Move()
			case "speak":
				selectedAnimal.Speak()
			default:
				fmt.Println("Unknown locomotion!")
			}
		} else {
			fmt.Println("Unknown action!")
		}
	}
}
