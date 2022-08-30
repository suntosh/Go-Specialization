package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var vo float64
	var so float64
	var t float64

	//var displacement float64

	var fn func(float64) float64

	fmt.Println("Please enter the aceleration, initial velocity and initial displacement (separate by space):")
	fmt.Scanln(&a, &vo, &so)
	fmt.Println("Please enter the time:")
	fmt.Scanln(&t)

	fn = GenDisplaceFn(a, vo, so)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

func GenDisplaceFn(a float64, vo float64, so float64) func(float64) float64 {

	fn := func(t float64) float64 {
		fmt.Println("Calculating the displacement:")
		return (a/2)*math.Pow(t, 2) + (vo * t) + so

	}
	return fn
}
