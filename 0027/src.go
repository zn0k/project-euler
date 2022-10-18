package main

import (
	"fmt"
	"math"
)

func sieve(n int) []int {
	sieve := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		sieve[i] = true
	}
	sieve[0], sieve[1] = false, false

	prime := 2
	for prime*prime <= n {
		for i := 2; i*prime < n; i++ {
			sieve[i*prime] = false
		}
		for i := prime + 1; i < n; i++ {
			if sieve[i] {
				prime = i
				break
			}
		}
	}

	var primes []int
	for i := 0; i < n; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func check(a, b int, p *map[int]bool) int {
	count := 0
	for n := 0; ; n++ {
		x := n*n + a*n + b
		_, ok := (*p)[x]
		if ok {
			count++
		} else {
			return count
		}
	}
}

func main() {
	primes := make(map[int]bool)
	for _, p := range sieve(100_000) {
		primes[p] = true
	}

	longest := 0
	best_a, best_b := 0, 0

	for a := -999; math.Abs(float64(a)) < 1000; a++ {
		for b := -1000; math.Abs(float64(b)) <= 1000; b++ {
			ps := check(a, b, &primes)
			if ps > longest {
				longest = ps
				best_a = a
				best_b = b
			}
		}
	}
	fmt.Printf("%d\n", best_a*best_b)
	fmt.Printf("a=%d, b=%d produces %d primes\n", best_a, best_b, longest)
}
