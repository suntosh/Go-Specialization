package main

import (
	"fmt"
	"os"
	s "strings"
)

func exit(w string) {
	fmt.Println("Not Found!")
	os.Exit(0)
}

func main() {

	var word string
	fmt.Println("Enter the word to be printed")
	_, err := fmt.Scanf("%s", &word)
	if err != nil {
		fmt.Println("Invalid Operation")
		return
	}

	wlen := len(word) - 1
	if s.Index(word, "i") == 0 || s.Index(word, "I") == 0 {
		if s.LastIndex(word, "n") == wlen || s.LastIndex(word, "N") == wlen {
			if s.IndexAny(word, "a") > -1 || s.IndexAny(word, "A") > -1 {
				fmt.Println("\nFound!")
			} else {
				exit("a")
			}
		} else {
			exit("n")
		}
	} else {
		exit("i")
	}

}
