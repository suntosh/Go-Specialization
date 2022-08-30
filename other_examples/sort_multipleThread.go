package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func sortTask(arr []int, wg *sync.WaitGroup) {
	count := len(arr)
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if arr[j] < arr[i] {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			}
		}
	}
	wg.Done()
}

func mergeArrays(arr1 []int, arr2 []int, out []int) {
	size1 := len(arr1)
	size2 := len(arr2)
	i := 0
	j := 0
	k := 0
	value := 0
	for i < size1 && j < size2 {
		if arr1[i] < arr2[j] {
			value = arr1[i]
			i++
		} else {
			value = arr2[j]
			j++
		}
		out[k] = value
		k++
	}

	for ; i < size1; i++ {
		out[k] = arr1[i]
		k++
	}

	for ; j < size2; j++ {
		out[k] = arr2[j]
		k++
	}
}

func sortArray(arr []int) []int {
	size := len(arr)

	// Handle small size < 4
	if size < 4 {
		var wg sync.WaitGroup
		wg.Add(1)
		sortTask(arr, &wg)
		wg.Wait()
		return arr
	}

	k := size / 4

	var slice1 []int = arr[0:k]
	var slice2 []int = arr[k : 2*k]
	var slice3 []int = arr[2*k : 3*k]
	var slice4 []int = arr[3*k : size]

	fmt.Println("Subarray1:")
	fmt.Println(slice1)

	fmt.Println("Subarray2:")
	fmt.Println(slice2)

	fmt.Println("Subarray3:")
	fmt.Println(slice3)

	fmt.Println("Subarray4:")
	fmt.Println(slice4)

	var wg sync.WaitGroup
	wg.Add(4)
	go sortTask(slice1, &wg)
	go sortTask(slice2, &wg)
	go sortTask(slice3, &wg)
	go sortTask(slice4, &wg)
	wg.Wait()

	// Merge 4 subarrays into one sorted array
	out1 := make([]int, len(slice1)+len(slice2))
	mergeArrays(slice1, slice2, out1)

	out2 := make([]int, len(slice3)+len(slice4))
	mergeArrays(slice3, slice4, out2)

	out := make([]int, len(out1)+len(out2))
	mergeArrays(out1, out2, out)

	return out
}

func main() {

	fmt.Printf("Enter a sequence of integers separated by spaces: ")
	// bufio.NewScanner will allow for spaces
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sli := strings.Fields(scanner.Text())

	arr := make([]int, len(sli))
	for i, s := range sli {
		arr[i], _ = strconv.Atoi(s)
	}

	fmt.Println("Input array before sorted:")
	fmt.Println(arr)

	out := sortArray(arr)

	fmt.Println("Output array after sorted:")
	fmt.Println(out)
}
