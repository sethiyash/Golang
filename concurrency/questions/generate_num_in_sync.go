// Problem is that you have the empty slice now create two functions addOdd and addEven
// that should run in goroutine and make sure to append the odd/even number in the orders
// such that array will have elements from 1 to 20 in sorted order
package main

import (
	"fmt"
	"sync"
)

func addOddNumbers(nums *[]int, oddChan, evenChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 20; i += 2 {
		<-oddChan
		*nums = append(*nums, i)
		evenChan <- true
	}
}

func addEvenNumber(nums *[]int, oddChan, evenChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 20; i += 2 {
		<-evenChan
		*nums = append(*nums, i)
		if i != 20 {
			oddChan <- true
		}
	}
}

func main() {
	nums := []int{}
	oddChan := make(chan bool)
	evenChan := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)

	// refer that the slices are pass by reference but when you pass them in go routine it
	// creates its own copy and that why we have to pass the address and use it as pointer in the methods
	go addOddNumbers(&nums, oddChan, evenChan, &wg)
	go addEvenNumber(&nums, oddChan, evenChan, &wg)

	oddChan <- true

	fmt.Println(nums)

	defer close(evenChan)
	defer close(oddChan)
}

