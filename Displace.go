/*=============================================================================
|   Assignment: Week 2 Assignment 1
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
|        Input: 	4 floats
|        Output:	1 float
|		 Algorithm: fn(s) = ½ a t^2 + V * t + s
| 		 Required Features Not Included:  None
|   	 Known Bugs:  None
|===========================================================================*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	values := make([]float64, 3)
	accel := true
	velocity := true
	displace := true
	valuefor := "Acceleration"
	fmt.Printf("Enter value for %s:\n", valuefor)
	fmt.Printf("> ")
	for reader.Scan() {
		val, err := strconv.ParseFloat(reader.Text(), 10)
		if err != nil {
			fmt.Printf("Oops! that was'nt taken well - try different\n")
			fmt.Printf("> ")
			continue
		}
		if accel {
			values[0] = val
			accel = false
			valuefor = "Velocity"
		} else if velocity {
			values[1] = val
			velocity = false
			valuefor = " Initial Displacement"

		} else if displace {
			values[2] = val
			displace = false
		}
		if !displace {
			break
		}
		fmt.Printf("Enter value for %s:\n", valuefor)
		fmt.Printf("> ")
	}
	fn := GetDisplaceFn(values[0], values[1], values[2])
	fmt.Println("Enter a value for time or enter 'x' to exit :")
	fmt.Printf("> ")
	for reader.Scan() {
		if strings.EqualFold(reader.Text(), "x") {
			os.Exit(0) //we exit
		}
		val, err := strconv.ParseFloat(reader.Text(), 10)
		if err != nil {
			fmt.Printf("Oops! that was'nt taken well - try different\n")
			fmt.Printf("> ")
			continue
		}
		fmt.Printf("This displacement for time %f is %f\n", val, fn(val))
		fmt.Printf("> ")
	}
}

func GetDisplaceFn(val ...float64) func(float64) float64 {
	return func(time float64) float64 {
		//s =  ½ a t2 + vo * t + s
		return .5*val[0]*math.Pow(time, 2) + val[1]*time + val[2]
	}
}
