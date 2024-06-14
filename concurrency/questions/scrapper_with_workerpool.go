//Implement a worker pool in Go to manage concurrent tasks such as web scraping. 
// you should create a simple worker pool to handle a dummy scraping operation for a given list of URLs

package main

import (
	"fmt"
	"sync"
	"time"
)

func DummyScrapper(url string) {
	fmt.Println("Started scrapping: ", url)
	time.Sleep(1 * time.Second)
	fmt.Println("Finished Scrapping: ", url)
}

// Worker function should receive URLs from channel and process them using dummy scrapper
func Worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	for url := range jobs {
		DummyScrapper(url)
		wg.Done()
	}
}

// Dispatcher sets the worker pool and distributes URL to the workers
func Dispatcher(urls []string, numWorkers int) {
	jobs := make(chan string, len(urls))
	var wg sync.WaitGroup

	for id := 1; id <= numWorkers; id++ {
		go Worker(id, jobs, &wg)
	}

	for _, url := range urls {
		wg.Add(1)
		jobs <- url
	}

	close(jobs)

	wg.Wait()
}

func main() {
	urls := []string{"https://www.google.com", "https://www.google.com/1"}

	numWorkers := 5

	// starting Dispatcher
	Dispatcher(urls, numWorkers)
}
