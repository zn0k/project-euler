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

func combinations(n int) []int {
	var result []int
	ds := digits(n)
	l := len(ds)
	for i := 0; i < l; i++ {
		r := 0
		for j, k := range ds {
			r += k * int((math.Pow10(l - j - 1)))
		}
		result = append(result, r)
		ds = append(ds[1:], ds[0])
	}
	return result
}

func check(p int, ps map[int]bool) bool {
	allPrime := true
	cs := combinations(p)
	for _, c := range cs {
		_, ok := ps[c]
		if !ok {
			allPrime = false
			break
		}
	}
	return allPrime
}

func main() {
	primes := make(map[int]bool)
	for _, p := range sieve(1_000_000) {
		primes[p] = true
	}

	total := 0
	for p := range primes {
		if check(p, primes) {
			total++
		}
	}
	fmt.Printf("%d\n", total)
}
