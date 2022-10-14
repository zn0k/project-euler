package main

import "fmt"

func count(n int) int {
	result := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			result += 2
			if n/i == i {
				result--
			}
		}
	}
	return result
}

func factorCount(n int) int {
	c := 0
	i := 1
	for ; c <= n; i++ {
		if i%2 == 0 {
			c = count(i/2) * count(i+1)
		} else {
			c = count(i) * count((i+1)/2)
		}
	}
	return i - 1
}

func main() {
	i := factorCount(500)
	fmt.Printf("%d\n", i*(i+1)/2)
}
