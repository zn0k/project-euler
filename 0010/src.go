package main

import (
	"fmt"
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

func main() {
	primes := sieve(2000000)
	sum := 0
	for _, p := range primes {
		sum += p
	}
	fmt.Printf("%d\n", sum)
}
