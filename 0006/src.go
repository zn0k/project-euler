package main

import "fmt"

func main() {
	sumOfSquares := 0
	squareOfSums := 0
	for i := 1; i <= 100; i++ {
		sumOfSquares += (i * i)
		squareOfSums += i
	}
	squareOfSums *= squareOfSums
	fmt.Printf("%d\n", squareOfSums-sumOfSquares)
}
