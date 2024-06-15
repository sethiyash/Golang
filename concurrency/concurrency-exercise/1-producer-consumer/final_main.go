//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, tweetsChan chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}
		tweetsChan <- tweet
	}
	close(tweetsChan)
	return
}

func consumer(tweetsChan <-chan *Tweet) {
	for t := range tweetsChan {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	var wg sync.WaitGroup

	// Producer
	tweetsChan := make(chan *Tweet)

	wg.Add(2)
	// we can also have a waitGroup to indicate producer is finished
	go func() {
		producer(stream, tweetsChan)
		wg.Done()
	}()

	// Consumer
	// we can create a anonymous func here and mark wg.Done() otherwise we have to pass wg in the original
	// function as parameter
	go func() {
		consumer(tweetsChan)
		wg.Done()
	}()

	wg.Wait()
	// if we are putting consumer in go routine we have to wait for consumer to finish we can achieve it with
	// 1. sync.waitGroup
	// 2. we can assume that the consumer will finish once the channel get closed

	// Approach 2
	// for range tweetsChan {
	// }

	fmt.Printf("Process took %s\n", time.Since(start))
}
