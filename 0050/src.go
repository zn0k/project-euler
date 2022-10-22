package main

import "fmt"

const LIMIT = 1_000_000

func Sieve(n int) []int {
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
	ps := Sieve(LIMIT)
	lookup := make(map[int]bool)
	for _, p := range ps {
		lookup[p] = true
	}

	// the sequence is a least 21 primes long, and the total is under one million. once we get over (1_000_000/21),
	// we only need to look at 21 more items, or any longer window further right would be over a million
	// find the first prime over (1_000_000/21)
	max_i := 0
	for _, p := range ps {
		if p > LIMIT/21 {
			break
		}
		max_i++
	}
	max_i += 21

	// sum primes as they go
	sums := make([]int, max_i+1)
	sums[0] = 0
	for i := 1; i < max_i+1; i++ {
		sums[i] = sums[i-1] + ps[i-1]
	}
	sum_length := len(sums)

	// start with big windows and make them smaller, sliding them over the primes
	// each window has a sum of its last element minus its first element
done:
	for length := sum_length - 2; length > 0; length-- {
		for start := 1; start+length < sum_length; start++ {
			sum := sums[start+length] - sums[start-1]
			if _, ok := lookup[sum]; ok {
				fmt.Printf("%d is composed of %d primes starting at %d\n", sum, length, ps[start])
				break done
			}
		}
	}
}
