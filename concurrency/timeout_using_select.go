// the time.After function returns a channel that blocks for the specified duration 
// after the duration the channel delivers the current time, once
// run on - 
package main

import (
	"fmt"
	"time"
)

type message struct {
	str  string
}

func boring(m message) <-chan message {
	c := make(chan message, 1)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(100 * time.Millisecond)
			c <- message{fmt.Sprintf("%s: %d", m.str, i)}
		}
	}()
	return c
}

func main() {
	timeOut := time.After(2*time.Second)
	c := boring(message{"Joe"})
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("You are too slow")
			return
		case s := <-c:
			fmt.Println(s.str)
		case <-timeOut:
			fmt.Println("Complete timeOut")
			return
		}
	}
}