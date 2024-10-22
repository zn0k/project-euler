package main

import "fmt"

const LIMIT = 10

func main() {
	nums := make([][]int, LIMIT+1)

	for n := 2; n <= LIMIT; n += 1 {
		for i := n + n; i <= LIMIT; i += n {
			nums[i] = append(nums[i], n)
		}
	}

	for i := 2; i <= LIMIT; i += 1 {
		phi := len(nums[i]) + 1
		totient := float64(i) / float64(phi)
		fmt.Printf("n = %d, relative primes = %v, phi = %d, totient = %f\n", i, nums[i], phi, totient)
	}
}
