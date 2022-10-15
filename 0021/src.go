package main

import (
	"fmt"
)

func sumOfDivisors(n int) int {
	sum := 0
	for i := 1; i <= (n/2)+1; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	nums := make(map[int]int)

	for i := 1; i < 10000; i++ {
		nums[i] = sumOfDivisors(i)
	}

	sum := 0
	for k, v := range nums {
		if nums[v] == k && v != k {
			sum += k
		}
	}
	fmt.Printf("%d\n", sum)
}
