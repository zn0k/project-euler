package main

import "fmt"

func main() {
	x, y := 1, 2
	sum := 2

	for y < 4000000 {
		x, y = y, x+y
		if y%2 == 0 {
			sum += y
		}
	}
	fmt.Printf("%d\n", sum)
}
