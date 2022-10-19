package main

import (
	"fmt"
	"math"
)

const LIMIT = 100_000_000

func pentagonal(max int) []int {
	var result []int

	pen := 0
	for n := 1; pen < max; n++ {
		pen = (n * (3*n - 1)) / 2
		result = append(result, pen)
	}
	return result
}

func check(i, j int, lookup map[int]bool) bool {
	_, ok := lookup[i+j]
	if !ok {
		return false
	}
	_, ok = lookup[j-i]
	if !ok {
		return false
	}
	return true
}

func main() {
	ps := pentagonal(LIMIT)
	lookup := make(map[int]bool)
	for _, p := range ps {
		lookup[p] = true
	}

	dist := LIMIT
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			if check(ps[i], ps[j], lookup) {
				d := math.Abs(float64(ps[j]) - float64(ps[i]))
				if int(d) < dist {
					dist = int(d)
				}
			}
		}
	}
	fmt.Println(dist)
}
