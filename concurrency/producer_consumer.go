package main

import "fmt"

func producer(c chan<- int, done chan<- bool) {
	for i:=0; i<10; i++ {
		c <- i
	}
	done <- true
}

func consumer(c <-chan int) {
	for  {
		fmt.Printf("Consumed: %d\n", <-c)
	}
}


func main() {
	c := make(chan int)
	done := make(chan bool)

	go producer(c, done)
	go consumer(c)

	<-done
}
