package main

import "fmt"

func isAbundant(n int) bool {
	sum := 0
	for i := 1; i <= (n/2)+1; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return (sum > n)
}

func main() {
	upperBound := 28123
	lowerBound := 12

	var abundants []int

	for i := lowerBound; i <= upperBound-lowerBound; i++ {
		if isAbundant(i) {
			abundants = append(abundants, i)
		}
	}

	lookup := make(map[int]bool)
	for _, i := range abundants {
		for _, j := range abundants {
			sum := i + j
			if sum <= upperBound {
				lookup[sum] = true
			}
		}
	}

	sum := 0
	var solution []int
	for target := 1; target <= upperBound; target++ {
		_, ok := lookup[target]
		if !ok {
			sum += target
			solution = append(solution, target)
		}
	}
	fmt.Printf("%d\n", sum)
}
