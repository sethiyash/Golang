package main

import (
	"fmt"
)

type message struct {
	str  string
	wait chan bool
}

func boring(m message) <-chan message {
	c := make(chan message)
	go func() {
		for i := 0; ; i++ {
			c <- message{fmt.Sprintf("%s: %d", m.str, i), m.wait}
			<-m.wait
		}
	}()

	//time.Sleep(100 * time.Millisecond)
	return c
}

func fanIn(input1, input2 <-chan message) <-chan message {
	c := make(chan message)
	go func() {
		for {
			// using select statement to handle multiple channels
			// its like a switch but each case is communication
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := fanIn(boring(message{"Ann", make(chan bool)}), boring(message{"Joe", make(chan bool)}))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg2.wait <- true
		msg1.wait <- true
	}

}
