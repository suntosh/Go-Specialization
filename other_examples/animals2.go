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

	// not specified in the assignment, but makes life easier
	getName() string
}

type Cow struct {
	name string
}

func (cow *Cow) Eat() {
	fmt.Printf("grass")
}

func (cow *Cow) Move() {
	fmt.Printf("walk")
}

func (cow *Cow) Speak() {
	fmt.Printf("moo")
}

func (cow *Cow) getName() string {
	return cow.name
}

type Bird struct {
	name string
}

func (bird *Bird) Eat() {
	fmt.Printf("worms")
}

func (bird *Bird) Move() {
	fmt.Printf("fly")
}

func (bird *Bird) Speak() {
	fmt.Printf("peep")
}

func (bird *Bird) getName() string {
	return bird.name
}

type Snake struct {
	name string
}

func (snake *Snake) Eat() {
	fmt.Printf("mice")
}

func (snake *Snake) Move() {
	fmt.Printf("slither")
}

func (snake *Snake) Speak() {
	fmt.Printf("hsss")
}

func (snake *Snake) getName() string {
	return snake.name
}

func main() {

	animals := make([]Animal, 0)

	for {

		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)
		var input string
		if scanner.Scan() {
			input = scanner.Text()
		}

		splitInput := strings.SplitN(input, " ", -1)

		if len(splitInput) != 3 {
			fmt.Printf("Error. Too many / too few arguments \n")
		}

		command := splitInput[0]
		firstParam := splitInput[1]
		secondParam := splitInput[2]

		// fmt.Scan(&command, &firstParam, &secondParam)
		command = strings.ToLower(command)
		secondParam = strings.ToLower(secondParam)

		if command == "newanimal" {
			var newAnimal Animal

			switch secondParam {
			case "cow":
				newAnimal = &Cow{firstParam}
			case "snake":
				newAnimal = &Snake{firstParam}
			case "bird":
				newAnimal = &Bird{firstParam}
			default:
				newAnimal = nil
			}

			if newAnimal != nil {
				animals = append(animals, newAnimal)
				fmt.Printf("“Created it!” \n")
			}
		} else if command == "query" {

			for _, animal := range animals {
				if animal.getName() == firstParam {
					switch secondParam {
					case "eat":
						animal.Eat()

					case "move":
						animal.Move()

					case "speak":
						animal.Speak()

					default:
						fmt.Printf("command not recognized.")
					}

					fmt.Printf("\n")
				}
			}
		}

	}
}
