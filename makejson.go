/*=============================================================================
|   Assignment: Week 4 Assignment 1
|   Author:  	Santosh Ahuja
|   Language:  	Go Version 1.16 on Windows 10
|   To Compile: go run makejason.go
|
|   Class:  	Getting Started with Go
|   Instructor: Ian Harris
|   Due Date:   April 16th
|
+-----------------------------------------------------------------------------
|
|  Description:  Input name and address as strings and outputs a json .
|
|        Input: 	string pairs
|        Output:	Jsonified name and adddresses
|		 Algorithm: Infinite loop
| 		 Required Features Not Included:  None
|   	 Known Bugs:  None
*===========================================================================*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
 * Main functions that runs the entire routine
 */

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Enter name amd then enter address : ")
	idMap := make(map[string]string)
	key := true
	var name, addr string
	errorstring := "Name"
	fmt.Print("> ")
	for scanner.Scan() {
		if scanner.Text() == "" {
			fmt.Printf("Enter something for %s\n", errorstring)
			fmt.Print("> ")
			continue
		}
		if key {
			name = scanner.Text()
			idMap["name"] = name
			errorstring = "Address"
			key = false
		} else {
			addr = scanner.Text()
			idMap["address"] = addr
			jsonStr, _ := json.Marshal(idMap)
			fmt.Println(string(jsonStr))
			key = true
			break
		}
		fmt.Print("> ")
	}
}
