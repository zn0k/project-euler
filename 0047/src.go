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

func primeFactors(n int, ps []int) []int {
	var fs []int
	for _, p := range ps {
		if p > n {
			break
		}
		for n%p == 0 {
			// normally, we should append p multiple times if needed
			// however, for this problem, 2 twice counts as 2**2, and one factor
			// so only count each p once
			if len(fs) == 0 || fs[len(fs)-1] != p {
				fs = append(fs, p)
			}
			n /= p
		}
		if n == 1 {
			break
		}
	}
	return fs
}

func main() {
	ps := sieve(LIMIT)

	fs := make([][]int, LIMIT)
	fs[644] = primeFactors(644, ps)
	fs[645] = primeFactors(645, ps)
	fs[646] = primeFactors(646, ps)
	for i := 646; i < LIMIT; i++ {
		fs[i] = primeFactors(i, ps)
		if len(fs[i]) == 4 && len(fs[i-1]) == 4 && len(fs[i-2]) == 4 && len(fs[i-3]) == 4 {
			fmt.Println(i - 3)
			break
		}
	}
}
