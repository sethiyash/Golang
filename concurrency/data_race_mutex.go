// Here Data races occur when we increment counter variable without sufficient protections
// leading to undefined or unpredictable behavior.

package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu sync.Mutex
)

func incrementCounter() {
	for i:=0; i<1000; i++ {
		mu.Lock()
		counter++  // this is shared data so need to acquire a lock before writing into it
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	for i:=0; i<10; i++ { // spanning 10 goroutines concurrently
		wg.Add(1)   // adding waitgroup to wait for them before the main goroutine exists
		fmt.Println("Running Goroutine: ", i)
		go func() {
			defer wg.Done()
			incrementCounter()
		}()
	}
	wg.Wait() // will wait until all the goroutines are not finished
	fmt.Printf("counter Value: %d\n", counter)
}
