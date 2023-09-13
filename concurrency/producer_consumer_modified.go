package main

import (
"context"
"fmt"
"math"
"time"
)

//Producer
//* It shall generate an incremental number every 700 milliseconds.
//* It shall exit on its own after 4 seconds
//* Consumer
//* It shall print generated number
//* It shall exit after printing all numbers

type Producer struct {
	num  *chan int
}

type Consumer struct {
	num *chan int
}

func (p *Producer) produce(ctx context.Context) error {
	for i := 0; i < math.MaxInt; i++ {
		time.Sleep(700 * time.Millisecond)
		*p.num <- i
	}
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("Producer has stopped after 4 seconds")
		}
	}
}

func (c *Consumer) consume() {
	for {
		number := <-*c.num
		fmt.Printf("Consumed number: %d\n", number)
	}
}

func server() (int, error) {
	var num = make(chan int)
	//var done = make(chan bool)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	producer := Producer{num: &num}
	go func() {
		producer.produce(ctx)
	}()

	consumer := Consumer{num: &num}
	go consumer.consume()

	for {
		select {
		case <-ctx.Done():
			//*producer.done <- true
			return 666, fmt.Errorf("program has timed out after 4 seconds")
		}
	}
}

func main() {
	start := time.Now()
	res, err := server()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	fmt.Println("Total time: ", time.Since(start))
}
