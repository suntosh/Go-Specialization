/*=============================================================================
|   Assignment: Week 2 Assignment 1
|   Author:  	Santosh Ahuja
|   Language:  	Go Version 1.16 on Windows 10
|   To Compile: go run Read.go path_to_file_name
|
|   Class:  	Concurrency In Go
|   Instructor: Ian Harris
|   Due Date:   April 16th
|
+-----------------------------------------------------------------------------
|
|  Description:  Race conditions occuring between Goroutines accessing shared data.
|  Examplanation : There are two routines Doughnuts maker and Doughnut eater. Doughnut maker
| 				 makes doughnuts at interval of 1 second and Doughnuteater routine eats
|				 doughnuts at 500 miliseond per second so there is a race condition taking
|				 taking the number of doughnuts to negative as the doughnuts are being eaten
|				 faster than they are made. Everytime the program is run the value of the
|				 available donughnuts is unpredictable and inconsistent with time.
|
|===========================================================================*/

package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

//Doughnuts
var doughnuts int = 0

//Waitgroup to sync completion of routines
var wg sync.WaitGroup

//Doughnut maker
func makeDoughnuts() {
	wg.Add(1)
	defer wg.Done()
	for {
		time.Sleep(1 * time.Second)
		doughnuts++
		fmt.Printf("Doughnuts Available (Made ---> :|) %d\n", doughnuts)
		Exit()
	}

}

//Doughnut eater
func eatDoughnut() {
	wg.Add(1)
	defer wg.Done()
	for {
		time.Sleep(799 * time.Millisecond)
		doughnuts--
		fmt.Printf("Doughnuts Available (Eaten <--- :O) %d\n", doughnuts)
		Exit()
	}
}

//
func main() {
	go makeDoughnuts()
	time.Sleep(5 * time.Second) // Lets make some dough nuts before they get eaten
	go eatDoughnut()
	wg.Wait()
}

func Exit() {
	if doughnuts < 0 {
		fmt.Println("We have eaten doughnuts faster than they are made")
		os.Exit(0)
	}
}
