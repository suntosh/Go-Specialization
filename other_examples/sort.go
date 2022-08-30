package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partitions := make([][]int, 4)
	sliceIntegers := make([]int, 0)
	sortedSliceIntegers := make([]int, 0)
	c := make(chan []int, 1)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter integers separated by a space: ")
	input, _ := reader.ReadString('\n')
	arrStrings := strings.Split(strings.TrimSpace(input), " ")

	for _, v := range arrStrings {
		val, _ := strconv.Atoi(v)
		sliceIntegers = append(sliceIntegers, val)
	}
	generatePartitions(sliceIntegers, &partitions) // generates the partitions based on the total of entered integers

	for i := 0; i < 4; i++ {
		go sortSlice(partitions[i], c, i+1) //start goroutines that will sort the smaller slices
	}

	for i := 0; i < 4; i++ {
		sortedSlice := <-c                                                // read the data from the channel (sorted slices)
		sortedSliceIntegers = append(sortedSliceIntegers, sortedSlice...) // concatenate the different sorted slices
	}

	sortBigSlice(sortedSliceIntegers) //sort and print the final array
}

func generatePartitions(numbers []int, partitions *[][]int) {
	for i := 0; i < len(numbers); i++ {
		(*partitions)[i%4] = append(((*partitions)[i%4]), numbers[i])
	}

	for i, x := range *partitions {
		fmt.Printf("Partition %d: ", i+1)
		for _, y := range x {
			fmt.Printf("%d ", y)
		}
		fmt.Println()
	}

}

func sortSlice(partition []int, c chan []int, id int) {
	sli := partition
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(sli)-1; i++ {
			if sli[i] > sli[i+1] {
				temp := sli[i+1]
				sli[i+1] = sli[i]
				sli[i] = temp
				swap = true
			}
		}
	}
	fmt.Println("GoRoutine ", id, "sorting:", partition)
	c <- sli //write result into channel
}

func sortBigSlice(numbers []int) {
	sli := numbers
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(sli)-1; i++ {
			if sli[i] > sli[i+1] {
				temp := sli[i+1]
				sli[i+1] = sli[i]
				sli[i] = temp
				swap = true
			}
		}
	}
	fmt.Printf("Main final sorting: ")

	for _, x := range sli {
		fmt.Printf("%d ", x)
	}
	fmt.Println()

}
