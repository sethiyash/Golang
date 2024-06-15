package main

import (
	"fmt"
	"sync"
)

// Problem Statement:
// You need to implement a basic Pub/Sub system in Go where:

// A publisher sends messages to a topic.
// Subscribers receive messages from the topic they subscribed to.

// Subscribe: Adds a subscriber to a topic.
// Unsubscribe: Removes a subscriber from a topic.
// Publish: Sends messages to all subscribers of a topic.

type Subscriber chan string

type Topic struct {
	subscribers map[Subscriber]bool
	mu          sync.RWMutex
}

func NewTopic() *Topic {
	return &Topic{
		subscribers: make(map[Subscriber]bool),
	}
}

func (t *Topic) Subscribe() Subscriber {
	sub := make(Subscriber)
	t.mu.Lock()
	t.subscribers[sub] = true
	t.mu.Unlock()
	return sub
}

func (t *Topic) Unsubscribe(sub Subscriber) {
	t.mu.Lock()
	if _, found := t.subscribers[sub]; found {
		delete(t.subscribers, sub)
		close(sub)
	}
	t.mu.Unlock()
}

func (t *Topic) Publish(message string) {
	t.mu.RLock()
	for sub := range t.subscribers {
		sub <- message
	}
	t.mu.RUnlock()
}

type Publisher struct {
	topics map[string]*Topic
	mu     sync.RWMutex
}

func NewPublisher() *Publisher {
	return &Publisher{
		topics: make(map[string]*Topic),
	}
}

func main() {
	publisher := NewPublisher()
	topic1 := publisher.topics["topic1"]
	sub1 := topic1.Subscribe()
	sub2 := topic1.Subscribe()

	topic2 := publisher.topics["topic2"]
	sub3 := topic2.Subscribe()

	go func() {
		for msg := range sub1 {
			fmt.Println("Subscriber 1:", msg)
		}
	}()

	go func() {
		for msg := range sub2 {
			fmt.Println("Subscriber 2:", msg)
		}
	}()

	go func() {
		for msg := range sub3 {
			fmt.Println("Subscriber 3:", msg)
		}
	}()

	topic1.Publish("Hello Topic 1")
	topic2.Publish("Hello Topic 2")

	topic1.Unsubscribe(sub2)

	topic1.Publish("Bye Topic 1")

	select {}
}
