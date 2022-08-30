package main 

import (
	"fmt"
	"sync"
)

type Chpstk struct {
	sync.Mutex
}

type Philo struct {
	id      int
	meals   int
	chpstks map[*Chpstk]struct{}
}

func NewPhilo(id int, lChpstk, rChpstk *Chpstk) *Philo {
	return &Philo{
		id:      id,
		meals:   0,
		chpstks: map[*Chpstk]struct{} {lChpstk:{}, rChpstk:{}},
	}
}

func (p *Philo) Eat() {
	// Maps do not guarantee an iteration order
	for c := range p.chpstks {
		c.Lock()
	}
	fmt.Printf("starting to eat %v\n", p.id)
	p.meals++
	fmt.Printf("finishing eating %v\n", p.id)
	for c := range p.chpstks {
		c.Unlock()
	}
}

func (p *Philo) Dine(permChan chan bool, wg *sync.WaitGroup) {
	for {
		// Wait for permission
		select {
		case <-permChan:
			p.Eat()
			wg.Done()
		}
	}
}

type Host struct {
	philos     map[*Philo]chan bool
	maxMeals   int
	maxConcEat int
}

func NewHost(philos []*Philo, maxMeals, maxConcEat int) *Host {
	newHost := Host{
		philos:     map[*Philo]chan bool{},
		maxMeals:   maxMeals,
		maxConcEat: maxConcEat,
	}

	for _, p := range philos {
		newHost.philos[p] = make(chan bool)
	}

	return &newHost
}

func (host *Host) HostMeal(done chan struct{}) {
	wg := sync.WaitGroup{}
	// Start the philosopher waiting goroutines
	for p := range host.philos {
		go p.Dine(host.philos[p], &wg)
	}

	// List of currently chosen philosophers
	eatingPhilos := make([]*Philo, 0, host.maxConcEat)
	for len(host.philos) > 0 {
		// Choose philosophers to eat
		for p := range host.philos {
			eatingPhilos = append(eatingPhilos, p)
			if len(eatingPhilos) == cap(eatingPhilos) {
				break
			}
		}
		fmt.Println() // To differentiate eating groups in log
		// Give permission
		wg.Add(len(eatingPhilos))
		for _, p := range eatingPhilos {
			host.philos[p] <- true
		}
		// Wait for them to finish
		wg.Wait()
		for _, p := range eatingPhilos {
			if p.meals >= host.maxMeals {
				delete(host.philos, p) // If the philosopher has eaten enough meals, remove him from the map
			}
		}
		eatingPhilos = eatingPhilos[:0]
	}
	done <- struct{}{}
}

func main() {
	// Make 5 chopsticks
	chpstks := make([]Chpstk, 5)
	for i := range chpstks {
		chpstks[i] = Chpstk{}
	}
	// Make 5 philosophers, sharing chopsticks
	philos := make([]*Philo, 5)
	for i := range philos {
		philos[i] = NewPhilo(i, &chpstks[i], &chpstks[(i+1)%5])
	}
	// Make the host
	host := NewHost(philos, 3, 2)

	done := make(chan struct{})
	// Start the meal, philosopher goroutines are started within
	go host.HostMeal(done)
	// Block until host signals the meal is done
	<- done
}
