package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name, food, locomotion, sound string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}
func (a *Animal) Move() {
	fmt.Println(a.locomotion)

}
func (a *Animal) Speak() {
	fmt.Println(a.sound)

}

var snake Animal = Animal{"Snake", "Mice", "Slither", "hssssss"}
var cow Animal = Animal{"Cow", "Grass", "Walk", "mooooo"}
var bird Animal = Animal{"Bird", "Worms", "Fly", "peeeep"}

func main() {
	var target Animal
	input := ""
	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command [eg. 'animal action'] >")
		sc.Scan()
		input = sc.Text()
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "exit" {
			break
		}
		commands := strings.Split(input, " ")

		if len(commands) < 2 {
			fmt.Println("Invalid command. Expecting 2 arguments")
			continue
		}

		switch commands[0] {
		case "snake":
			target = snake
		case "cow":
			target = cow
		case "bird":
			target = bird
		default:
			fmt.Printf("Animal '%s' not Found!\n", commands[0])
		}

		switch commands[1] {
		case "speak":
			target.Speak()
		case "eat":
			target.Eat()
		case "move":
			target.Move()
		default:
			fmt.Printf("Command '%s' not found!\n", commands[1])
		}

	}

}
