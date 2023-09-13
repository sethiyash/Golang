package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
The classical Dining philosophers problem.

Implemented with forks (aka chopsticks) as mutexes.
*/

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
on a line by itself, where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
on a line by itself, where <number> is the number of the philosopher.
*/

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

func(p Philosopher) dineModified() {
	defer eatWgroup.Done()

	// as each philosopher can eat 3 times
	for i:=0; i<3; i++ {
		say("Thinking", p.id)
		randomPause(2)

		say("Hungry", p.id)
		p.leftFork.Lock()
		p.rightFork.Lock()

		say("Eating", p.id)
		randomPause(1)
		p.leftFork.Unlock()
		p.rightFork.Unlock()
	}
}

var eatWgroup sync.WaitGroup

func main() {
	// number of philosophers
	count := 5

	forks := make([]*Fork, count)
	for i:=0; i<count; i++ {
		forks[i]=new(Fork)
	}

	philosophes := make([]*Philosopher, count)
	for i:=0; i<count; i++{
		philosophes[i] = &Philosopher{i, forks[i], forks[(i+1)%count]}
		eatWgroup.Add(1)
		go philosophes[i].dineModified()
	}
	eatWgroup.Wait()
}