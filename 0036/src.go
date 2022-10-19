package main

import (
	"fmt"
)

func digits(n int) []int {
	var result []int

	for n > 10 {
		result = append(result, n%10)
		n /= 10
	}
	result = append(result, n)
	return result
}

func toBin(n int) []int {
	var result []int
	for n > 0 {
		result = append(result, n&1)
		n = n >> 1
	}
	return result
}

func palindrome(ns []int) bool {
	for i := 0; i < len(ns)/2; i++ {
		j := len(ns) - i - 1
		if ns[i] != ns[j] {
			return false
		}
	}
	return true
}

func main() {
	sum := 0
	for i := 0; i < 1_000_000; i++ {
		if palindrome(digits(i)) && palindrome(toBin(i)) {
			sum += i
		}
	}
	fmt.Printf("%d\n", sum)
}
