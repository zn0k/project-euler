package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat/combin"
)

const LIMIT int = 20000

func Cube(n int) int {
	return n * n * n
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

func main() {
	var cubes []int
	lookup := make(map[int]bool)
	for i := int(0); i <= LIMIT; i++ {
		n := Cube(i)
		cubes = append(cubes, n)
		lookup[n] = true
	}

	for _, i := range cubes {
		ps := Permutations(i, lookup)
		if len(ps) == 5 {
			fmt.Println(i)
			break
		}
	}
}
