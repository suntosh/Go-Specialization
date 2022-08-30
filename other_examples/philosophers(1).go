package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Chopstick struct {
	mux sync.Mutex
}

type Philosopher struct {
	number int
	eatCounter int
	leftChopstick *Chopstick
	rightChopstick *Chopstick
}

func (p *Philosopher) getChopsticks() {
	if rand.Int() % 2 == 1 {
		fmt.Printf("Philosopher %v picked left chopstick\n", p.number)
		p.leftChopstick.mux.Lock()
		fmt.Printf("Philosopher %v picked right chopstick\n", p.number)
		p.rightChopstick.mux.Lock()
	} else {
		fmt.Printf("Philosopher %v picked right chopstick\n", p.number)
		p.rightChopstick.mux.Lock()
		fmt.Printf("Philosopher %v picked left chopstick\n", p.number)
		p.leftChopstick.mux.Lock()
	}
}

func (p *Philosopher) returnChopsticks() {
	if rand.Int() % 2 == 1 {
		fmt.Printf("Philosopher %v returned left chopstick\n", p.number)
		p.leftChopstick.mux.Unlock()
		fmt.Printf("Philosopher %v returned right chopstick\n", p.number)
		p.rightChopstick.mux.Unlock()
	} else {
		fmt.Printf("Philosopher %v returned right chopstick\n", p.number)
		p.rightChopstick.mux.Unlock()
		fmt.Printf("Philosopher %v returned left chopstick\n", p.number)
		p.leftChopstick.mux.Unlock()
	}
}

func (p *Philosopher) eat() {
	p.getChopsticks()
	fmt.Printf("starting to eat %v\n", p.number)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("finishing eating %v\n", p.number)
	p.returnChopsticks()
}

func (p *Philosopher) dine(feast *sync.WaitGroup, host *Host) {
	fmt.Printf("Philosopher %v starts dining\n", p.number)
	for ;p.eatCounter < 3; {
		if host.requestPermission() {
			fmt.Printf("Philosopher %v was granted permission to eat\n", p.number)
			p.eat()
			p.eatCounter++
			host.doneEating()
		} else {
			fmt.Printf("Philosopher %v was denied permission to eat\n", p.number)
		}
		time.Sleep(100 * time.Millisecond)
	}
	feast.Done()
}

func (h *Host) requestPermission() bool {
	h.mux.Lock()
	defer h.mux.Unlock()
	if h.queue < 2 {
		h.queue++
		return true
	}
	return false
}

func (h *Host) doneEating() {
	h.mux.Lock()
	defer h.mux.Unlock()
	h.queue--
	return
}

func prepareFeast() ([]Philosopher, *Host) {
	chopsticks := make([]Chopstick,5,5)
	philosophers := make([]Philosopher,0,5)
	for i := 1; i <= 5; i++ {
		philosophers = append(philosophers,Philosopher{i,0,&chopsticks[i-1],&chopsticks[i%5]})
		fmt.Printf("Added philosopher %v\n", i)
	}
	var host Host
	return philosophers, &host
}

type Host struct {
	mux sync.Mutex
	queue int
}

func main() {
	philosophers, host := prepareFeast()
	var feast sync.WaitGroup
	for i, _ := range(philosophers) {
		go philosophers[i].dine(&feast, host)
	}
	feast.Add(5)
	feast.Wait()
}
