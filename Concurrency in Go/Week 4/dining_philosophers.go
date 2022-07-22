package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id                     int
	lChopstick, rChopstick *Chopstick
}

var host = make(chan bool, 2) // host channel

func (p *Philosopher) Eat(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		host <- true // this will block if host channel already have 2 values

		p.lChopstick.Lock()
		p.rChopstick.Lock()

		fmt.Printf("Starting to eat %d\n", p.id)
		time.Sleep(time.Second) // assume eating takes 1 second
		fmt.Printf("Finishing eating %d\n", p.id)

		p.lChopstick.Unlock()
		p.rChopstick.Unlock()

		<-host
	}

	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}

	// Create 5 chopsticks
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	// Create 5 philosophers and assign their left and right chopstick
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			i,
			chopsticks[i],
			chopsticks[(i+1)%5],
		}
	}

	for _, p := range philosophers {
		wg.Add(1)
		go p.Eat(&wg)
	}

	wg.Wait()
}
