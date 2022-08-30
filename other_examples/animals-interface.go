package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (cow *Cow) Eat() {
	fmt.Println("grass")
}

func (cow *Cow) Move() {
	fmt.Println("walk")
}

func (cow *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (bird *Bird) Eat() {
	fmt.Println("worms")
}

func (bird *Bird) Move() {
	fmt.Println("fly")
}

func (bird *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

func (snake *Snake) Eat() {
	fmt.Println("mice")
}

func (snake *Snake) Move() {
	fmt.Println("slither")
}

func (snake *Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	animalMap := make(map[string]Animal)

	for {
		fmt.Print("> ")
		var strCommand string
		var strName string
		var arg3 string
		fmt.Scanf("%s %s %s", &strCommand, &strName, &arg3)

		switch strCommand {

		case "newanimal":
			switch arg3 {
			case "cow":
				animalMap[strName] = &Cow{strName}
			case "bird":
				animalMap[strName] = &Bird{strName}
			case "snake":
				animalMap[strName] = &Snake{strName}
			}
			fmt.Println("Created it!")

		case "query":
			switch arg3 {
			case "eat":
				animalMap[strName].Eat()
			case "move":
				animalMap[strName].Move()
			case "speak":
				animalMap[strName].Speak()
			}
		}
	}
}
