/* 
	Write a Bubble Sort program in Go. The program
	should prompt the user to type in a sequence of up to 10 integers. The program
	should print the integers out on one line, in sorted order, from least to
	greatest. Use your favorite search tool to find a description of how the bubble
	sort algorithm works.

*/
package main

import "fmt"

func scanSlice(n int) ([]int, error) {
    fmt.Printf("Please Input %d numbers: ", n)
    x := make([]int, n)
    y := make([]interface{}, len(x))
    for i := range x {
        y[i] = &x[i]
    }
    n, err := fmt.Scanln(y...)
    x = x[:n]
    return x, err
}

func BubbleSort(numbers []int) {
	for i := len(numbers); i > 0; i-- {
	   for j := 1; j < i; j++ {
		  if numbers[j-1] > numbers[j] {
				Swap(numbers, j-1)
			}
		}
	}
}

func Swap(arr []int, index int) {
	temp := arr[index]
	arr[index] = arr[index + 1]
	arr[index + 1] = temp
}

func main() {
	var count int
	fmt.Printf("How many number would you like to sort? ")
	fmt.Scanln(&count)
    x, err := scanSlice(count)
    if err != nil {
        fmt.Println(err)
        return
    }
	BubbleSort(x)
	fmt.Printf("Bubble Sort Result: ")
	fmt.Println(x)
}
