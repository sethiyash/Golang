// implement thread safe map in GO
// we can implement it using sync.Mutex
package main

import (
    "fmt"
    "hash/fnv"
    "sync"
)

type SafeMap struct {
    mu    sync.Mutex
    store map[string]int
}

func NewSafeMap() *SafeMap {
    return &SafeMap{store: make(map[string]int)}
}

func (m *SafeMap) Set(key string, value int) {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.store[key] = value
}

func (m *SafeMap) Get(key string) (int, bool) {
    m.mu.Lock()
    defer m.mu.Unlock()
    val, ok := m.store[key]
    return val, ok
}

func main() {
    sm := NewSafeMap()
    sm.Set("a", 1)
    value, exists := sm.Get("a")
    if exists {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Value not found")
    }
}

// Q: What is a potential drawback of using a single Mutex for a highly concurrent map?
// Using a single Mutex for a highly concurrent map can lead to contention, 
// where multiple goroutines are blocked waiting to acquire the lock. 
// This can become a bottleneck and degrade performance, especially in scenarios with high read/write throughput.

 // Q: How can you improve the performance of a thread-safe map under high contention?
//  One way to improve performance is by using sharded maps. 
// By dividing the map into several shards, each protected by its own Mutex, 
// you can reduce contention and allow multiple goroutines to access different shards concurrently.

const numShards = 32

type ShardedMap struct {
    shards [numShards]*SafeMap
}

func NewShardedMap() *ShardedMap {
    m := &ShardedMap{}
    for i := 0; i < numShards; i++ {
        m.shards[i] = NewSafeMap()
    }
    return m
}

func (m *ShardedMap) getShard(key string) *SafeMap {
    h := fnv.New32a()
    h.Write([]byte(key))
    return m.shards[h.Sum32()%numShards]
}

func (m *ShardedMap) Set(key string, value int) {
    shard := m.getShard(key)
    shard.Set(key, value)
}

func (m *ShardedMap) Get(key string) (int, bool) {
    shard := m.getShard(key)
    return shard.Get(key)
}

 // Q: What is the difference between sync.Mutex and sync.RWMutex?
// sync.Mutex allows only one goroutine to access the critical section at a time, whether for reading or writing. 
// sync.RWMutex, on the other hand, distinguishes between read and write locks. 
// Multiple goroutines can acquire the read lock simultaneously, allowing concurrent reads, but the write lock is exclusive. 
// RWMutex provides better performance in scenarios with a high read-to-write ratio.




