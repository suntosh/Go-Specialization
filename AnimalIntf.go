/*=============================================================================
|   Assignment: Week 4 Assignment1
|   Author:  	Santosh Ahuja
|   Language:  	Go Version 1.16 on Windows 10
|   To Compile: go run Read.go path_to_file_name
|
|   Class:  	Methods and Interfaces In Go
|   Instructor: Ian Harris
|   Due Date:   April 16th
|
+-----------------------------------------------------------------------------
|
|  Description:  GetDisplaceFn example
|
|        Input: 	commands "newanimal" or "query"
|        Output:	string
|		 Algorithm: hashmaps
| 		 Required Features Not Included:  None
|   	 Known Bugs:  None
|===========================================================================*/

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

type cow struct {
}

func (cow) Eat() {
	fmt.Println("grass")
}
func (cow) Move() {
	fmt.Println("walk")
}
func (cow) Speak() {
	fmt.Println("moo")
}

type bird struct {
}

func (bird) Eat() {
	fmt.Println("worms")
}
func (bird) Move() {
	fmt.Println("fly")
}
func (bird) Speak() {
	fmt.Println("peep")
}

type snake struct {
}

func (snake) Eat() {
	fmt.Println("mice")
}
func (snake) Move() {
	fmt.Println("slither")
}
func (snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter \"newanimal\" OR \"query\" followed by two parameters :")
	fmt.Printf("> ")
	for reader.Scan() {
		command := strings.ToLower(reader.Text())
		if strings.HasPrefix(command, "newanimal") {
			params := strings.Fields(command)
			if len(params) == 3 {
				switch params[2] {
				case "cow":
					animals[params[1]] = cow{}
				case "bird":
					animals[params[1]] = bird{}
				case "snake":
					animals[params[1]] = snake{}
				default:
					{
						fmt.Printf("Invalid animal type : \"%s\", Try again\n", params[2])
						fmt.Printf("> ")
						continue
					}
				}
			} else {
				fmt.Println("Please enter EXACTLY TWO parameters for \"newanimal\" command")
				fmt.Println("> ")
				continue
			}

		} else if strings.HasPrefix(command, "query") {
			params := strings.Fields(command)
			if len(params) == 3 {
				var animIntf Animal
				animIntf = animals[params[1]]
				if animIntf == nil {
					fmt.Printf("\"%s\" animal not found\n", params[1])
					fmt.Printf("> ")
					continue
				}
				switch params[2] {
				case "eat":
					animIntf.Eat()
				case "move":
					animIntf.Move()
				case "speak":
					animIntf.Speak()
				default:
					{
						fmt.Printf("Invalid query: \"%s\"\n", params[2])
						fmt.Printf("> ")
						continue
					}
				}
			} else {
				fmt.Println("Please enter EXACTLY TWO parameters for \"newanimal\" command")
				fmt.Printf("> ")
				continue
			}
		} else {
			fmt.Println("Invalid command Please enter \"newanimal\" OR \"query\" followed by two parameters")
			fmt.Printf("> ")
			continue
		}
		fmt.Printf("> ")
	}

}
