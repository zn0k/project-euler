package main

import "fmt"

func factoradic(seed string, n int) []int {
	l := len(seed)
	n -= 1

	result := make([]int, l)
	for i := 1; i <= l; i++ {
		result[l-i] = n % i
		n /= i
	}
	return result
}

func nth(seed string, n int) string {
	var chars []string
	for _, r := range seed {
		chars = append(chars, string(r))
	}

	result := ""
	for _, i := range factoradic(seed, n) {
		result += chars[i]
		chars = append(chars[:i], chars[i+1:]...)
	}
	return result
}

func main() {
	fmt.Printf("%s\n", nth("0123456789", 1_000_000))
}
