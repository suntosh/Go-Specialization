/*=============================================================================
|   Assignment: Week 3 Assignment 1
|   Author:  	Santosh Ahuja
|   Language:  	Go Version 1.16 on Windows 10
|   To Compile: go run slice.go
|
|   Class:  	Getting Started with Go
|   Instructor: Ian Harris
|   Due Date:   April 16th
|
+-----------------------------------------------------------------------------
|
|  Description:  Inpput nt arbitrary number of integers and populate it in an
|				array that grows with input. Array is sorted at every input.
|
|        Input: 	integer .
|        Output:	Prints the sorted array with input integers
|		 Algorithm: Recursive
| 		 Required Features Not Included:  None
|   	 Known Bugs:  None
*===========================================================================*/

package main

/* Required Imports */
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*Global variable scanner variable */
var scanner = bufio.NewScanner(os.Stdin)

/*
 *
 * This is the main function that has two loops
 * 1st loop populates the array slice with initial size 3
 * The second call is to a recursive function that keeps
 * populating array with integers and expands as needed
 *
 */

func main() {
	sint := make([]int, 3)
	fmt.Println("Enter an integer to store OR 'x' to exit:")
	var val int
	for i := 0; i < 3; {
		if acceptIntOrX(&val) {
			sint[0] = val
			sort.Ints(sint)
			i++
			fmt.Println(sint)
		}
	}
	fillIntArray(sint)
}

/*
 * @Functiobn fillIntArray - Recusrsive function to fill and expand array.
 * @var sint - int array
 */
func fillIntArray(sint []int) {
	var y int
	if acceptIntOrX(&y) {
		sint = append(sint, y)
		sort.Ints(sint)
		fmt.Println(sint)
	}
	fillIntArray(sint)
}

/*
 * @Functiobn acceptIntorX- Functon that processes user input
 * @var x  reference that containes that processed input integer
 * @return bool - success or not at processing user input
 */
func acceptIntOrX(x *int) bool {
	var input string
	scanner.Scan()
	input = scanner.Text()
	if strings.EqualFold(input, "x") {
		os.Exit(0) //we exit
	} else {
		i, err := strconv.ParseInt(input, 10, 64)
		if handleError(err) {
			return false
		} else {
			*x = int(i) // but we will just round it to int for KISS
			return true
		}
	}
	return false
}

/*
 * @Function handleError - handles error if occured
 * @var error - error stack
 * @retrun bool - if there is an error condition or not.
 */
func handleError(err error) bool {
	if err != nil {
		fmt.Println("Does'nt seem an integer or 'x' - Try again")
		return true
	}
	return false
}
