package main

import "fmt"

const LIMIT = 10_000_000_000

func triangles(max int) []int {
	var result []int
	t := 0
	for n := 1; t < max; n++ {
		t = (n * (n + 1)) / 2
		result = append(result, t)
	}
	return result
}

func pentagonal(max int) map[int]bool {
	result := make(map[int]bool)
	t := 0
	for n := 1; t < max; n++ {
		t = (n * (3*n - 1)) / 2
		result[t] = true
	}
	return result
}

func hexagonal(max int) map[int]bool {
	result := make(map[int]bool)
	t := 0
	for n := 1; t < max; n++ {
		t = n * (2*n - 1)
		result[t] = true
	}
	return result
}

func main() {
	ts := triangles(LIMIT)
	ps := pentagonal(LIMIT)
	hs := hexagonal(LIMIT)
	for i := 285; i < len(ts); i++ {
		_, ok := ps[ts[i]]
		if !ok {
			continue
		}
		_, ok = hs[ts[i]]
		if !ok {
			continue
		}
		fmt.Println(ts[i])
		break
	}
}
