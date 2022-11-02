package main

import (
	"fmt"
	"strconv"
)

func triangle(n int) int {
	return (n * (n + 1)) / 2
}

func square(n int) int {
	return n * n
}

func pentagonal(n int) int {
	return (n * (3*n - 1)) / 2
}

func hexagonal(n int) int {
	return n * (2*n - 1)
}

func heptagonal(n int) int {
	return (n * (5*n - 3)) / 2
}

func octagonal(n int) int {
	return n * (3*n - 2)
}

func generate(f func(int) int, lower int, upper int) map[string]bool {
	result := make(map[string]bool)
	for i := 1; ; i++ {
		n := f(i)
		if n < lower {
			continue
		}
		if n >= upper {
			break
		}
		s := strconv.FormatInt(int64(n), 10)
		result[s] = true
	}
	return result
}

func split(s string) (string, string) {
	return s[:2], s[2:]
}

func main() {
	//functions := []func(int) int{triangle, square, pentagonal, hexagonal, heptagonal, octagonal}
	functions := []func(int) int{triangle, square, pentagonal}
	for _, f := range functions {
		fmt.Println(generate(f, 1000, 10000))
	}
}
