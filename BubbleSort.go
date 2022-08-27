/*=============================================================================
|   Assignment: Week 1 Assignment 1
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
|  Description:  BubbleSort example
|
|        Input: 	10 integers
|        Output:	Sorted 10 intgers
|		 Algorithm: BubbleSort
| 		 Required Features Not Included:  None
|   	 Known Bugs:  None
|===========================================================================*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	intArray := make([]int64, 10)
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("=============Program BubbleSort============")
	fmt.Println("Enter 10 integers to sort :")
	fmt.Print("> ")
	counter := 0
	for reader.Scan() {
		i, err := strconv.ParseInt(reader.Text(), 10, 64)
		if err != nil {
			fmt.Printf("Oops! that was'nt taken well - try different\n")
			fmt.Printf("> ")
			continue
		}
		intArray[counter] = i
		if counter == len(intArray)-1 {
			break
		} else {
			counter++
		}
		fmt.Printf("> ")
	}
	BubbleSort(intArray)
	fmt.Printf("The Sorted Array Is %d", intArray)
}

/* Bubble sort - iteratively compares and swap places with the neighbourhood */
func BubbleSort(sli []int64) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(sli)-1; i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
				sorted = false
			}
		}
	}
}

/* Elegant looking Swap function Go, Go! */
func Swap(sli []int64, x int) {
	sli[x+1], sli[x] = sli[x], sli[x+1]
}
