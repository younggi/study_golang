package main

import "fmt"

type IntSet map[int]int

func NewIntSet() IntSet {
	return IntSet{}
}

func (m IntSet) Insert(val int) {
	m[val]++
}

func (m IntSet) Erase(val int) {
	if _, exists := m[val]; !exists {
		return
	}
	m[val]--
	if m[val] <= 0 {
		delete(m, val)
	}
}

func (m IntSet) Count(val int) int {
	return m[val]
}

func (m IntSet) String() string {
	vals := ""
	for val, count := range m {
		for i := 0; i < count; i++ {
			vals += fmt.Sprint(val) + " "
		}
	}
	return "{ " + vals + "}"
}
