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

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	reader := bufio.NewScanner(os.Stdin)
	var animal *Animal
	fmt.Println("Enter animal( cow | bird |snake ) <space> properties( eat | move | speak ) :")
	fmt.Printf("> ")
	for reader.Scan() {
		names := strings.Fields(reader.Text())
		if len(names) != 2 {
			fmt.Println("Oops! that was'nt taken well - try TWO words ONLY seperated by space")
			fmt.Printf("> ")
			continue
		}
		switch strings.ToLower(names[0]) {
		case "cow":
			animal = &cow
		case "bird":
			animal = &bird
		case "snake":
			animal = &snake
		default:
			{
				fmt.Printf("your input %s is invalid option\n", names[0])
				fmt.Printf("> ")
				continue
			}
		}

		switch strings.ToLower(names[1]) {
		case "eat":
			fmt.Println(animal.Eat())
		case "move":
			fmt.Println(animal.Move())
		case "speak":
			fmt.Println(animal.Speak())
		default:
			{
				fmt.Printf("your input %s is invalid option\n", names[1])
				fmt.Printf("> ")
				continue
			}
		}
		fmt.Printf("> ")
	}
}
