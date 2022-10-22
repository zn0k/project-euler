package main

import (
	"fmt"
	"math"
)

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

func Digits(n int) []int {
	var result []int

	for n >= 10 {
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

func Combine(ds []int) int {
	r := 0
	l := len(ds)
	for j, k := range ds {
		r += k * int((math.Pow10(l - j - 1)))
	}
	return r
}

func replace(xs []int, x, y int) int {
	result := make([]int, len(xs))
	for i := 0; i < len(xs); i++ {
		if xs[i] == x {
			result[i] = y
		} else {
			result[i] = xs[i]
		}
	}

	combined := Combine(result)
	return combined
}

func count(n int, lookup map[int]bool) int {
	ds := Digits(n)
	max := 1
	for i := 0; i < 10; i++ {
		matches := make(map[int]bool)
		for j := 0; j < 10; j++ {
			candidate := replace(ds, i, j)
			if _, ok := lookup[candidate]; ok {
				cs := Digits(candidate)
				if len(ds) == len(cs) {
					matches[candidate] = true
				}
			}
		}
		if len(matches) > max {
			max = len(matches)
		}
	}

	return max
}

func main() {
	ps := Sieve(LIMIT)
	lookup := make(map[int]bool)
	for _, p := range ps {
		lookup[p] = true
	}

	for _, p := range ps {
		if count(p, lookup) == 8 {
			fmt.Println(p)
			break
		}
	}
}
