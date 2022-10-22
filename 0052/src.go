package main

import (
	"fmt"
	"sort"
)

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

func cmp(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func check(n int) bool {
	ds := Digits(n)
	sort.Ints(ds)
	for m := 2; m <= 6; m++ {
		xs := Digits(n * m)
		sort.Ints(xs)
		if !cmp(ds, xs) {
			return false
		}
	}
	return true
}

func main() {
	x := 1
	for {
		if check(x) {
			fmt.Println(x)
			break
		}
		x++
	}
}
