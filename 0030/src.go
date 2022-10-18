package main

import (
	"fmt"
	"math"
)

func decompose(n int) []float64 {
	var result []float64

	for n >= 10 {
		result = append(result, float64(n%10))
		n /= 10
	}
	result = append(result, float64(n))
	return result
}

func check(n, e int) bool {
	var sum float64 = 0
	exp := float64(e)
	digits := decompose(n)

	for _, d := range digits {
		sum += math.Pow(d, exp)
	}
	return (sum == float64(n))
}

func main() {
	var solutions []int

	for i := 2; i < 500_000; i++ {
		if check(i, 5) {
			solutions = append(solutions, i)
		}
	}

	sum := 0
	for _, i := range solutions {
		sum += i
	}
	fmt.Printf("%d\n", sum)
}
