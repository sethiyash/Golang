// in 

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	var mu sync.Mutex
	order := 1
	n := 10 // number of goroutines

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				mu.Lock()
				if order == i {
					fmt.Printf("This is goroutine number %d\n", i)
					order++
					mu.Unlock()
					break
				}
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Total time to finish : %s \n", time.Since(start).String())
}

// Explanation
// WaitGroup: This ensures that the main function waits for all the goroutines to finish.
// Mutex (mu): This mutex is used to synchronize access to the shared variable order.
// Order Variable: This variable keeps track of the current order and is shared among all goroutines.
// Loop: Starts multiple goroutines, each representing a task.
// Goroutine Function: Each goroutine repeatedly locks the mutex, checks if it's their turn to print (by comparing with the order variable), prints the statement, increments the order, and then unlocks the mutex. If it's not their turn, they unlock the mutex and check again.
// Main Goroutine Waits: The main goroutine waits for all goroutines to complete using wg.Wait().
