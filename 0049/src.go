package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

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

func Unique(ns []int) []int {
	l := make(map[int]bool)
	var result []int
	for _, n := range ns {
		_, ok := l[n]
		if !ok {
			l[n] = true
			result = append(result, n)
		}
	}
	return result
}

func Permutations(n int, lookup map[int]bool) []int {
	var perms []int
	ds := Digits(n)
	iss := combin.Permutations(4, 4)
	for _, is := range iss {
		variant_ds := make([]int, 4)
		for i := 0; i < len(is); i++ {
			variant_ds[i] = ds[is[i]]
		}
		variant := Combine(variant_ds)
		_, ok := lookup[variant]
		if ok && variant >= 1000 {
			perms = append(perms, variant)
		}
	}
	return Unique(perms)
}

func Check(n int, lookup map[int]bool) (bool, string) {
	ps := Permutations(n, lookup)
	if len(ps) < 3 {
		return false, ""
	}
	iss := combin.Permutations(len(ps), 3)
	for _, is := range iss {
		terms := make([]int, 3)
		for i := 0; i < 3; i++ {
			terms[i] = ps[is[i]]
		}
		sort.Ints(terms)
		if terms[1]-terms[0] == terms[2]-terms[1] {
			s := strconv.FormatInt(int64(terms[0]), 10)
			s += strconv.FormatInt(int64(terms[1]), 10)
			s += strconv.FormatInt(int64(terms[2]), 10)
			return true, s
		}
	}
	return false, ""
}

func main() {
	ps := Sieve(10000)
	lookup := make(map[int]bool)
	for _, p := range ps {
		lookup[p] = true
	}

	for _, p := range ps {
		if p < 1000 || p > 9999 {
			continue
		}
		ok, result := Check(p, lookup)
		if ok && !strings.Contains(result, "1487") {
			fmt.Println(result)
			break
		}
	}
}
