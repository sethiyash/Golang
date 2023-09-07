// fan-in function or a multiplexing
// fan-in function allow to execute whosoever is ready
// Taken from - https://youtu.be/f6kdp27TYZs?t=866 (Go Concurrency Patterns Rob Pike)
// Run on - https://go.dev/play/p/87JRbljcWWH
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string { // Return receive only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d\n", msg, i)
		}
	}()
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	return c // return the channel to the caller
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are boring I'm leaving")
}


