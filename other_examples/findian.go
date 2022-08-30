package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Printf("Please enter a string: ")
	fmt.Scan(&input)

	findIAN(input)
}

func findIAN(input string) {
	lower := strings.ToLower(input)
	if strings.HasPrefix(lower, "i") && strings.HasSuffix(lower, "n") && strings.Contains(lower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}		
