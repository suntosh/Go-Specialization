package main

import (
	"fmt"
	"time"
)

func addandset(a int, c chan int) {
	for i := 0; i < 100; i++ {
		a++
		c <- a
		//time.Sleep(1 * time.Second)
	}
}

func main() {
	c := make(chan int)
	go addandset(2, c)
	time.Sleep(5 * time.Second)

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(<-c)
	}
}
