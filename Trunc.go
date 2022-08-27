package main

import "fmt"

func main() {
	var myfloat float64
	fmt.Println("Please enter your floating point ")
	_, err := fmt.Scanf("%f", &myfloat)
	if err != nil {
		fmt.Println("Invalid Syntax")
		return
	}

	fmt.Printf("The truncated float is %d\n", int(myfloat))

}
