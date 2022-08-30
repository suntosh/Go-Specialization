
/*==================================================================
|   Assignment: Week 4 Assignment 2
|   Author:  	Santosh Ahuja
|   Language:  	Go Version 1.16 on Windows 10
|   To Compile: go run Read.go path_to_file_name
|
|   Class:  	Concurrency In Go
|   Instructor: Ian Harris
|   Date:   	August 29th 2022|
+------------------------------------------------------------------
|
|  Description: Implemnent the Dining Philosophers Algo, 
|				Where synhronization is managed by a "Host"|  
|=================================================================*/
package main 

import (
	"fmt"
	"time"
	"sync"
)

type ChopS struct{ sync.Mutex  }

type Philo struct { leftCS , rightCS *ChopS 
					c chan int
					id int  }


func eat( p *Philo ) {
	for {
		<- p.c;
		fmt.Printf( "Start eating %d\n", p.id);
		p.leftCS.Lock();
		p.rightCS.Lock();
		time.Sleep(  300 * time.Millisecond);
		fmt.Println("Eating ", p.id);
		fmt.Printf( "Finishing eating %d\n", p.id);
		p.leftCS.Unlock();
		p.rightCS.Unlock();
	}
}


func main() {

	var chans = []chan int {
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
	 }

	fmt.Println("I am the host");
	CSticks := make([]*ChopS, 5)
	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
	   CSticks[i] = new(ChopS)
	}

	
	for i := 0; i < 5; i++ {
	   philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], chans[i] , i}
	}
	
	// start all the philosophers and get them to wait for a permission from the host
	go eat( philos[0]);
	go eat( philos[1]);
	go eat( philos[2]);
	go eat( philos[3]);
	go eat( philos[4]);
	

	for i := 0 ; i < 3 ; i++  {
		//host sends permission to 0 and 2 philosophers 
		chans[0] <- 1; 
		chans[2] <- 1;

		time.Sleep( 3 * time.Second); //lets them complete since only 2 are allowed 

		chans[1] <- 1; //host sends permission to 1 and 3 philosophers 
		chans[3] <- 1;

		time.Sleep( 3 * time.Second); //lets them complete since only 2 are allowed 

		chans[2] <- 1; // host sends permission to 2 and 4 philosophers 
		chans[4] <- 1;

		time.Sleep( 3 * time.Second); //lets them complete since only 2 are allowed 
	}

}
