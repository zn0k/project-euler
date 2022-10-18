package main

import "fmt"

func main() {
	dim := 1001
	sum := 1
	distance := 2
	start := 3

	for i := 0; i < (dim-1)/2; i++ {
		right_bottom := start
		left_bottom := right_bottom + distance
		left_top := left_bottom + distance
		right_top := left_top + distance
		sum += right_bottom + left_bottom + left_top + right_top
		start += 4*distance + 2
		distance += 2
	}
	fmt.Printf("%d\n", sum)
}
