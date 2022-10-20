package main

import "fmt"

const LIMIT = 1_000_000

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

func squares(n int) []int {
	var ss []int
	for i := 1; i*i < n; i++ {
		ss = append(ss, i*i)
	}
	return ss
}

func check(p int, ps []int, ss []int, lookup map[int]bool) bool {
	_, ok := lookup[p]
	if ok {
		return false
	}
	for i := 0; i < len(ps); i++ {
		if ps[i] > p {
			break
		}
		for j := 0; j < len(ss); j++ {
			test := ps[i] + 2*ss[j]
			if test == p {
				return false
			}
			if test > p {
				break
			}
		}
	}
	return true
}

func main() {
	ss := squares(LIMIT)
	ps := sieve(LIMIT)
	lookup := make(map[int]bool)
	for _, p := range sieve(LIMIT) {
		lookup[p] = true
	}

	for i := 9; i < LIMIT; i += 2 {
		if check(i, ps, ss, lookup) {
			fmt.Println(i)
			break
		}
	}
}
