package main

import (
	"fmt"
)

func main() {

	var floatNum float64
	_, err := fmt.Scan(&floatNum)
	if err != nil {
		fmt.Print("FAIL")
	} else {
		fmt.Println(int(floatNum))
	}
}
