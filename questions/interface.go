package main

import "fmt"

// Create two structs Slice and map and have following methods on them:
// addData(int)
// removeData(int)
// printData()

type CRUD interface {
	addData(int)
	removeData(int)
	printData()
}

type slice1 struct {
	data []int
}

type map1 struct {
	data map[int]int
}

// always remember when methods modify the receiver's data, so they should use pointer receivers
// Without this, the modifications will not affect the original data structures.
func (s *slice1) addData(val int) {
	s.data = append(s.data, val)
}

func (m *map1) addData(val int) {
	m.data[val] = 0
}

func (s *slice1) removeData(val int) {
	newSlice := make([]int, 0)
	for _, v := range s.data {
		if v != val {
			newSlice = append(newSlice, v)
		}
	}
	s.data = newSlice
}

func (m *map1) removeData(val int) {
	delete(m.data, val)
}

func (s *slice1) printData() {
	fmt.Println(s.data)
}

func (m *map1) printData() {
	fmt.Println(m.data)
}

func add(c CRUD, val int) {
	c.addData(val)
}

func remove(c CRUD, val int) {
	c.removeData(val)
}

func print(c CRUD) {
	c.printData()
}

func main() {
	s := &slice1{data: make([]int, 0)}
	m := &map1{data: make(map[int]int, 0)}

	add(s, 1)
	print(s)

	add(m, 1)
	print(m)
	remove(s, 1)
	print(s)

}
