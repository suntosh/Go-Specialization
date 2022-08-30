package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	length := 0
	sli := make([]int, length)
	symb := "X"
	fmt.Println("For exit enter 'X'")
	for length == 0 {
		intNum := ""
		fmt.Println("Please print interger: ")
		fmt.Scan(&intNum)
		if intNum == symb {
			break
		}
		test, err := strconv.Atoi(intNum)
		if err != nil {
			fmt.Println("Wrong number")
			continue
		}
		sli = (append(sli, test))
		sort.Sort(sort.IntSlice(sli))
		fmt.Println(sli)

	}
}
