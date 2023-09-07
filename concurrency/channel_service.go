// Generator: the function that returns a channel
// Taken from - https://youtu.be/f6kdp27TYZs?t=866 (Go Concurrency Patterns Rob Pike)
// Run on - https://go.dev/play/p/RdN1_DR913o
package main

import "fmt"

func boring(msg string) <-chan string { // Return receive only channel of strings
	c := make(chan string)
	go func() {
		for i:=0; ;i++ {
			c <- fmt.Sprintf("%s %d\n", msg, i)
		}
	}()
	return c // return the channel to the caller
}

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	for i:=0; i<5; i++ {  // here you will see both are maintaining the order
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You are boring I'm leaving")
}


