package main

import "fmt"

func factorial(n int) int {
	f := 1
	for i := n; i > 0; i-- {
		f *= i
	}
	return f
}

func digits(n int) []int {
	var result []int

	for n > 10 {
		result = append(result, n%10)
		n /= 10
	}
	result = append(result, n)
	return result
}

func main() {
	factorials := make(map[int]int)
	for i := 0; i < 10; i++ {
		factorials[i] = factorial(i)
	}

	total := 0
	for i := 10; i < 100000; i++ {
		sum := 0
		for _, d := range digits(i) {
			sum += factorials[d]
		}

		if sum == i {
			total += i
		}
	}
	fmt.Printf("%d\n", total)
}
