// Implement a custom map data structure in Go without using the built-in map 
// You can use chaining for collision resolution.

package main

import (
    "fmt"
    "hash/fnv"
)

type Entry struct {
    Key   string
    Value interface{}
    Next  *Entry
}

type HashMap struct {
    buckets []*Entry
    size    int
}

func NewHashMap(size int) *HashMap {
    return &HashMap{
        buckets: make([]*Entry, size),
        size:    size,
    }
}

func (h *HashMap) hash(key string) int {
    hasher := fnv.New32a()
    hasher.Write([]byte(key))
    return int(hasher.Sum32()) % h.size
}

func (h *HashMap) Put(key string, value interface{}) {
    index := h.hash(key)
    newEntry := &Entry{Key: key, Value: value}

    if h.buckets[index] == nil {
        h.buckets[index] = newEntry
    } else {
        current := h.buckets[index]
        for current.Next != nil {
            if current.Key == key {
                current.Value = value
                return
            }
            current = current.Next
        }
        if current.Key == key {
            current.Value = value
        } else {
            current.Next = newEntry
        }
    }
}

func (h *HashMap) Get(key string) (interface{}, bool) {
    index := h.hash(key)
    current := h.buckets[index]
    for current != nil {
        if current.Key == key {
            return current.Value, true
        }
        current = current.Next
    }
    return nil, false
}

func (h *HashMap) Delete(key string) {
    index := h.hash(key)
    current := h.buckets[index]
    var prev *Entry = nil

    for current != nil {
        if current.Key == key {
            if prev == nil {
                h.buckets[index] = current.Next
            } else {
                prev.Next = current.Next
            }
            return
        }
        prev = current
        current = current.Next
    }
}

func main() {
    hashMap := NewHashMap(10)
    hashMap.Put("name", "Alice")
    hashMap.Put("age", 30)

    if value, found := hashMap.Get("name"); found {
        fmt.Println("name:", value)
    }

    hashMap.Delete("name")

    if _, found := hashMap.Get("name"); !found {
        fmt.Println("name entry deleted")
    }
}
