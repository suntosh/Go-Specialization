/*=============================================================================
|   Assignment: Week 4 Assignment 2
|   Author:  	Santosh Ahuja
|   Language:  	Go Version 1.16 on Windows 10
|   To Compile: go run Read.go path_to_file_name
|
|   Class:  	Getting Started with Go
|   Instructor: Ian Harris
|   Due Date:   April 16th
|
+-----------------------------------------------------------------------------
|
|  Description:  Reads a file with space seperated first name and last name and
|               creates a slice of structs and then dumps the slice of structs
|
|        Input: 	File
|        Output:	pairs of first name and last name
|		 Algorithm: Infinite loop
| 		 Required Features Not Included:  None
|   	 Known Bugs:  None
|===========================================================================*/

package main

//imports
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* Person data type */

type Person struct {
	nameFirst string
	nameLast  string
}

const MAX_NAME_LENGTH = 20

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("\n\n")
		fmt.Println("       Usage: go run Read.go full_filename_with_path")
		fmt.Println()
		fmt.Println("       e.g  : go run Read.go \"D://test.txt\"")
		os.Exit(0)
	} else {
		readFile, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("File access error. Ensure file exists and the program can read")
			os.Exit(0)
		}
		Persons := make([]Person, 0)
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			var p Person
			var names []string = strings.Fields(fileScanner.Text())
			if len(names) > 1 { // we check for space errors in lines
				if len(names[0]) > MAX_NAME_LENGTH {
					p.nameFirst = names[0][:MAX_NAME_LENGTH]
				} else {
					p.nameFirst = names[0]
				}
			}
			if len(names) > 2 { // we check for space errors and empty names
				if len(names[1]) > MAX_NAME_LENGTH {
					p.nameLast = names[1][:MAX_NAME_LENGTH]
				} else {
					p.nameLast = names[1]
				}
			}
			Persons = append(Persons, p)
		}
		readFile.Close()
		for _, person := range Persons {
			fmt.Printf("First Name :%s ,Last Name :%s \n", person.nameFirst, person.nameLast)
		}
	}
}
