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

func variations(n int) []int {
	var result []int
	ds := digits(n)
	for i := 0; i < len(ds); i++ {
		r := 0
		s := ds[i:]
		l := len(s)
		for j, k := range s {
			r += k * int((math.Pow10(l - j - 1)))
		}
		result = append(result, r)

		r = 0
		s = ds[0 : len(ds)-i]
		l = len(s)
		for j, k := range s {
			r += k * int((math.Pow10(l - j - 1)))
		}
		result = append(result, r)
	}
	return result
}

func check(n int, ps map[int]bool) bool {
	if n < 10 {
		return false
	}
	for _, v := range variations(n) {
		_, ok := ps[v]
		if !ok {
			return false
		}
	}
	return true
}

func main() {
	primes := make(map[int]bool)

	for _, p := range sieve(1_000_000) {
		primes[p] = true
	}

	sum := 0
	for p := range primes {
		if check(p, primes) {
			sum += p
		}
	}
	fmt.Printf("%d\n", sum)
}
