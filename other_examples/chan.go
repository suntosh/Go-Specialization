package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
	eat_number      int
}

func (p Philo) eat(sit chan int, out chan int) {
	i := 0
	for {
		if i < p.eat_number {

			select {
			case msg := <-sit:
				//	fmt.Println("Philo SIT:", p.id)
				msg++
				rand.Seed(time.Now().UnixNano())
				cs_select := rand.Intn(2)

				//Randomly select right or left stick for lock. Random number is 1 or 0. cs_select = 1 or 0
				if cs_select == 1 {
					p.leftCS.Lock()
					p.rightCS.Lock()
					fmt.Printf("starting to eat %d\n", p.id)
					time.Sleep(300 * time.Millisecond)
					p.leftCS.Unlock()
					p.rightCS.Unlock()
					fmt.Printf("finishing eating %d\n", p.id)
					i++
					out <- p.id
				} else {
					p.rightCS.Lock()
					p.leftCS.Lock()
					fmt.Printf("starting to eat %d\n", p.id)
					time.Sleep(500 * time.Millisecond)
					p.rightCS.Unlock()
					p.leftCS.Unlock()
					fmt.Printf("finishing eating %d\n", p.id)
					i++
					out <- p.id
				}

			}

		} else {
			break
		}
	}
	//fmt.Println("Philo EXIT:", p.id)
	wg.Done()
}
func Host(sit chan int, out chan int, eat_number int) {
	sit <- 1
	sit <- 2

	i := 0
	counter := eat_number * 5

	for {
		if i < counter {
			select {
			case msg := <-out:
				//fmt.Println("Philo OUT:", msg)
				msg++
				i++
				if i < counter {
					sit <- i
				}

			default:
				if i == counter {
					close(sit)
				}
			}

		} else {
			break
		}

	}
	//	fmt.Println("HOST Exit")
	wg.Done()
}

func main() {
	sit := make(chan int, 2)
	out := make(chan int, 2)

	var eat_number int = 3

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5], eat_number}
	}

	wg.Add(1)
	go Host(sit, out, eat_number)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(sit, out)
	}

	wg.Wait()
}
