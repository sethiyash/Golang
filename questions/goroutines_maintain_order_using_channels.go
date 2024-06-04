// Make sure to print the numbers in correct order when print is called from a goroutine 

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	n := 3 // number of goroutines
	orderCh := make(chan int, 1)
	orderCh <- 1 // start with 1

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				// Receive the order number from the channel
				order := <-orderCh
				if order == i {
					fmt.Printf("This is goroutine number %d\n", i)
					// Send the next order number to the channel
					orderCh <- i + 1
					break
				}
				// Put back the order number to the channel
				orderCh <- order
			}
		}(i)
	}

	wg.Wait()
	close(orderCh)
}

// Explanation
// WaitGroup: This ensures that the main function waits for all the goroutines to finish.
// Channel (orderCh): This channel is used to coordinate the order of execution. It’s initialized with a buffer of 1 to hold the initial value.
// Initialization: The initial value (1) is sent to the channel to start the order.
// Loop: Starts multiple goroutines, each representing a task.
// Goroutine Function:
// Each goroutine receives the order number from the channel.
// - If the order number matches the goroutine’s number, it prints the statement and sends the next order number to the channel, then breaks the loop.
// - If the order number does not match, it puts the order number back into the channel and continues to wait for its turn.
// Main Goroutine Waits: The main goroutine waits for all goroutines to complete using wg.Wait() and then closes the channel.
