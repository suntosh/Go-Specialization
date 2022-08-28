
/*=============================================================================
|   Assignment: Week 3 Assignment 2
|   Author:  	Santosh Ahuja
|   Language:  	Go Version 1.16 on Windows 10
|   To Compile: go run Read.go path_to_file_name
|
|   Class:  	Concurrency In Go
|   Instructor: Ian Harris
|   Date:   	August 28th 2022
|
+-----------------------------------------------------------------------------
|
|  Description:   Taking 8 integers , splitting them in to 4 sub arrays , mergeing them into one array
|  |
|===========================================================================*/
package main 

import (
	"os"
	"fmt"
	"sync"
	"bufio"
	"strconv"
	"sort"
)

//Waitgroup to sync completion of routines
var wg sync.WaitGroup

func SortRoutine(arr []int, i int ) {
	wg.Add(1)
	defer wg.Done()
	fmt.Printf("Unsorted Array %d Is %v\n", i, arr);
	sort.Ints(arr);
	fmt.Printf("Sorted Array %d is %v\n", i, arr);
}


func main() {
	intArray := make([]int, 12);
	reader := bufio.NewScanner(os.Stdin) ;
	fmt.Println("Enter 12 integers to sort :") ;
	fmt.Print("> ") ;
	counter := 0 ;
	for reader.Scan() {
		i, err := strconv.ParseInt(reader.Text(), 10, 64)
		if err != nil {
			fmt.Printf("Oops! that was'nt taken well - try different\n")
			fmt.Printf("> ")
			continue
		}
		intArray[counter] = int(i)
		if counter == len(intArray)-1 {
			break
		} else {
			counter++
		}
		fmt.Printf("> ")
	}
	sub_array_1 := intArray[0:3];
	sub_array_2 := intArray[3:6];
	sub_array_3 := intArray[6:9];
	sub_array_4 := intArray[9:12];
	
	go SortRoutine ( sub_array_1 , 1);
	go SortRoutine ( sub_array_2, 2);
	go SortRoutine ( sub_array_3, 3);
	go SortRoutine ( sub_array_4, 4);

	wg.Wait(); 

	sortedSlice := append( sub_array_1, sub_array_2...) ;
	sortedSlice = append( sortedSlice, sub_array_3...);
	sortedSlice = append( sortedSlice, sub_array_4...);
	sort.Ints( sortedSlice);
	fmt.Printf("The Sorted Array Is %v\n", sortedSlice);
	
}
	
	