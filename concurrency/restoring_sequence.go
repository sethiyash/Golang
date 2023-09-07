// Restoring sequence that they need to be printed in the same order of how they are getting generated
// run it pn - https://go.dev/play/p/aO_pcnBPEs9
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type message struct {
	str  string
	wait chan bool
}

func boring(msg message) <-chan message { // Return receive only channel of strings
	c := make(chan message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- message{fmt.Sprintf("%s: %d", msg.str, i), waitForIt}
			<-waitForIt
		}
	}()
	time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
	return c // return the channel to the caller
}

func fanIn(input1, input2 <-chan message) <-chan message {
	c := make(chan message)
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
	c := fanIn(boring(message{"Joe", make(chan bool)}), boring(message{"Ann", make(chan bool)}))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You are boring I'm leaving")
}
