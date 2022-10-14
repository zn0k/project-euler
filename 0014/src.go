package main

import "fmt"

type Memoizer struct {
	m map[int]int
}

func NewMemoizer() *Memoizer {
	m := &Memoizer{}
	m.m = make(map[int]int)
	return m
}

func _collatz(n int) int {
	if n < 1 {
		panic("_collatz called with n < 1")
	}
	if n == 1 {
		return 1
	} else if n%2 == 0 {
		return 1 + collatz(n/2)
	} else {
		return 1 + collatz(n*3+1)
	}
}

func (m *Memoizer) Collatz(n int) int {
	i, ok := m.m[n]
	if ok {
		return i
	} else {
		val := _collatz(n)
		m.m[n] = val
		return val
	}
}

var m *Memoizer = NewMemoizer()

func collatz(n int) int { return m.Collatz(n) }

func main() {
	max := 0
	result := 0
	for i := 1; i < 1000000; i++ {
		val := collatz(i)
		if val > max {
			max = val
			result = i
		}
	}
	fmt.Printf("%d\n", result)
}
