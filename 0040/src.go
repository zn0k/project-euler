package main

import (
	"fmt"
)

func digits(n int) []int {
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

func fraction() chan int {
	c := make(chan int)
	num := 1
	go func() {
		for {
			ds := digits(num)
			for i := 0; i < len(ds); i++ {
				c <- ds[i]
			}
			num++
		}
	}()
	return c
}

func main() {
	trigger := 1
	product := 1
	count := 1
	for digit := range fraction() {
		if count == trigger {
			product *= digit
			if count == 1_000_000 {
				break
			}
			trigger *= 10
		}
		count++
	}
	fmt.Printf("%d\n", product)
}
