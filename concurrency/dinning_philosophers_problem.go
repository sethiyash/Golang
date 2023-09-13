/*
The classical Dining philosophers problem.
Implemented with forks (aka chopsticks) as mutexes.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Fork struct {
	sync.Mutex
}

type Philosopher struct {
	id int
	leftFork, rightFork *Fork
}

func say(msg string, id int) {
	fmt.Printf("%s: %d\n",msg, id)
}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

// Endlessly dine.
// Goes from thinking to hungry to eating and starts over.
// Adapt the pause values to increased or decrease contentions
// around the forks.
func (p Philosopher) dine() {
	say("thinking", p.id)
	randomPause(2)

	say("hungry", p.id)
	p.leftFork.Lock()
	p.rightFork.Lock()

	say("eating", p.id)
	randomPause(5)

	p.leftFork.Unlock()
	p.rightFork.Unlock()

	say("done", p.id)
	p.dine()
}


func main()  {
	// number of philosophers
	count := 5

	forks := make([]*Fork, count)
	for i:=0; i<count; i++ {
		forks[i] = new(Fork)
	}

	philosophes := make([]*Philosopher, count)
	for i:=0; i<count; i++{
		philosophes[i] = &Philosopher{i, forks[i], forks[(i+1)%count]}
		go philosophes[i].dine()
	}

	// wait endlessly while they are dinning
	endless := make(chan int)
	<-endless
}

