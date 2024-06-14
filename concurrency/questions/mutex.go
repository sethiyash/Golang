// Q: What is the zero value of a sync.Mutex in Go and how does it affect its operation?

// The zero value of a sync.Mutex is an unlocked mutex. In Go, zero values are usable without initialization, 
// which means a sync.Mutex can be used directly after declaration without explicitly initializing it. 
// This makes Mutex integration seamless and error-free in structures.

// Q: How does the sync.Map introduced in Go 1.9 differ from using a regular map with a Mutex?
// sync.Map is specifically designed for cases where the entry for a given key is only written once but read many times, 
// as in caches. sync.Map provides methods like Load, Store, and Delete which are safe to use concurrently. 
// Unlike a regular map with a Mutex, sync.Map optimizes read-heavy operations by reducing lock contention, 
// even though it might perform less efficiently in scenarios with frequent writes due to its internal complexity.

//Provide an example of a deadlock in Go using Mutex and explain how you would resolve it.

package main

import (
    "sync"
)

func main() {
    var mu sync.Mutex
    mu.Lock()
    mu.Lock() // This will deadlock, as the same goroutine tries to lock it again
}

// Resolution: To avoid this deadlock, ensure that a mutex is not locked again by the same goroutine without unlocking it first.
// This often involves more careful structuring of your code to prevent recursive locks or reentering a locked section.

// Q: Can you explain a scenario where replacing sync.Mutex with sync.RWMutex could significantly improve performance in a Go application?

// sync.RWMutex should be used in scenarios where data structure is read frequently but written to infrequently. 
// For instance, in a server application that holds a configuration map accessed by many requester goroutines to read configuration values and only
// occasionally updates the map values, using sync.RWMutex would allow these multiple readers to access the map concurrently without blocking each other, 
// thereby improving throughput and performance.

// Q: How would you implement a concurrency-safe LRU (Least Recently Used) cache in Go? What synchronization primitives would you use and why?

// Implementing a concurrency-safe LRU cache typically involves combining a doubly linked list with a hash map. The list is used to store the cache items in access order, and the map provides fast access to the nodes of the list so that items can be quickly moved to the front when accessed. sync.Mutex or sync.RWMutex can be used to synchronize access to the cache. sync.RWMutex is preferable because it allows multiple concurrent reads:

type LRUCache struct {
    capacity int
    cache    map[int]*list.Element
    list     *list.List
    lock     sync.RWMutex
}

// Implement methods to Get, Put which handle the synchronization with lock.RLock() for reads and lock.Lock() for writes

// Q: What problems might arise when using a Mutex inside a loop for controlling access to a map and how can you mitigate them?
// Using a Mutex inside a loop can lead to high contention and poor performance if the loop is executed very frequently or if the loop iterations are lengthy. 
// To mitigate this, you can reduce the granularity of the lock by trying to minimize the time the lock is held or by using finer-grained locking 
// (e.g., sharding the map and using multiple mutexes).

