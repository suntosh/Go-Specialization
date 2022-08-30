/*
The goal of this activity is to explore the use of threads by creating a 
program for sorting integers that uses four goroutines to create four 
sub-arrays and then merge the arrays into a single array.

Review criteria
Students will receive five points if the program sorts the integers and five 
additional points if there are four goroutines that each prints out a set of 
array elements that it is storing.
*/
package main

import (
    "fmt"
    "sort"
    "sync"
)

func fA(s []int, out chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
   
    fmt.Printf("function fA received: %v\n", s)
    sort.Ints(s)
    fmt.Printf("function fA sorted: %v\n", s)
    
    for _, x := range s {
        out <- x
    }
    close(out)
}

func fB(s []int, out chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
   
    fmt.Printf("function fB received: %v\n", s)
    sort.Ints(s)
    fmt.Printf("function fB sorted: %v\n", s)
    
    for _, x := range s {
        out <- x
    }
    close(out)
}

func fC(s []int, out chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
   
    fmt.Printf("function fC received: %v\n", s)
    sort.Ints(s)
    fmt.Printf("function fC sorted: %v\n", s)
    
    for _, x := range s {
        out <- x
    }
    close(out)
}

func fD(s []int, out chan<- int, wg *sync.WaitGroup) {
     defer wg.Done()
   
    fmt.Printf("function fD received: %v\n", s)
    sort.Ints(s)
    fmt.Printf("function fD sorted: %v\n", s)
    
    for _, x := range s {
        out <- x
    }
    close(out)
}

func merge(left []int, right []int)  []int {
    result := make([]int, len(left) + len(right))
      
    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            result[i] = left[0]
            left = left[1:]
        } else {
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }
      
    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }
      
    return result
}

func main() {
    // Get user input:
    var ival int
    var err error
    input := make([]int, 0)
    toEnter := 12
    
    for len(input) != toEnter {
        fmt.Printf("Enter an integer (%d to go): \n", toEnter - len(input))
        _, err = fmt.Scan(&ival)
        if err == nil {
            input = append(input, ival)
        } else {
            fmt.Printf("Error!\n")
        }
    }
    
    fmt.Printf("Values entered: %v\n", input)
    var wg sync.WaitGroup
    inA := make(chan int, 4)
    inB := make(chan int, 4)
    inC := make(chan int, 4)
    inD := make(chan int, 4)
    
    // Launch goroutines:
    wg.Add(1)
    go fA(input[:3], inA, &wg)
    
    wg.Add(1)
    go fB(input[3:6], inB, &wg)
    
    wg.Add(1)
    go fC(input[6:9], inC, &wg)

    wg.Add(1)
    go fD(input[9:], inD, &wg)

    // Wait for all goroutine to complete.
    wg.Wait()
    
    // Read in results from the buffered channels
    a := make([]int, 0)
    b := make([]int, 0)
    c := make([]int, 0)
    d := make([]int, 0)
    
    // Read in channel A:
    for v := range inA {
        a = append(a, v)
    }
    
    // Read in channel B:
    for v := range inB {
        b = append(b, v)
    }
    
    // Read in channel C:
    for v := range inC {
        c = append(c, v)
    }
    
    // Read in channel D:
    for v := range inD {
        d = append(d, v)
    }
    
    // Merge channels A and B:
    ab := merge(a, b)
    
    // Merge channels C and D:
    cd := merge(c, d)
    
    // Merge the final result:
    abcd := merge(ab, cd)
    
    fmt.Printf("sorted: %v\n", abcd)
}    
