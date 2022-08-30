package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v, d float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v*t + d
	}
}

func main() {
	var acceleration,
		velocity,
		displacement,
		time float64

	fmt.Printf("Enter the acceleration as a float: ")
	fmt.Scan(&acceleration)
	fmt.Printf("Enter the velocity as a float: ")
	fmt.Scan(&velocity)
	fmt.Printf("Enter the initial displacement as a float: ")
	fmt.Scan(&displacement)
	fmt.Printf("Enter the time as a float: ")
	fmt.Scan(&time)

	fmt.Printf("acceleration: %f, velocity: %f, initial displacement: %f\n time: %f\n", acceleration, velocity, displacement, time)

	fn := GenDisplaceFn(acceleration, velocity, displacement)

	fmt.Println(fn(time))
}
