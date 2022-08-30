package main

import (
	"fmt"
	"time"
)

func Add(x *int,y int) {
	*x = *x+1
	time.Sleep(1 * time.Second)
	fmt.Printf("\n%d,%d",*x,y)
	
}


func main() {
  var x int
  x = 1
  go Add(&x,1)
  //x = x+4
  go Add(&x,2)
  fmt.Printf("\n%d",x)
}
/*Race conditions is a problem that might happen as a result
		of all possible interleavings. The problem itself: outcome
		of the program depends on interleavings, which are not deterministic
		(interleaving of instructions between concurrent tasks is unknown,
		and depends on OS, Go runtime).

		Race conditions can occur due to communication between tasks (or goroutines),
		for example: multiple tasks are trying to access and manipulate the same shared
		variable, in this way those multiple tasks are communicating with each other
		through this shared variable. Due to the uncertainty of the Goroutine scheduling
		mechanism, the results of the program are unpredictable.


In this Add function using var x which used by both goroutine and simple function 
seems like before calling sleep funtion value of x is getting updated by both threads. this causes race condition for x.
*/