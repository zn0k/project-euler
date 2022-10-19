package main

import (
	"fmt"
	"sort"
)

func digits(n int) []int {
	var result []int

	for n > 10 {
		result = append(result, n%10)
		n /= 10
	}
	result = append(result, n)
	for i := 0; i < len(result)/2; i++ {
		j := len(result) - i - 1
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func pandigital(n int) bool {
	ds := digits(n)
	sort.Ints(ds)

	for i := 1; i <= len(ds); i++ {
		if ds[i-1] != i {
			return false
		} else {
		}
	}
	return true
}

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
	max := 0
	for _, p := range sieve(987654321) {
		if pandigital(p) && p > max {
			max = p
		}
	}
	fmt.Println(max)
}
