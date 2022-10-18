package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.NewInt(1)

	i := 2
	for {
		i++
		sum := a.Add(a, b)
		if len(sum.String()) == 1000 {
			fmt.Printf("%d\n", i)
			break
		}
		a = b
		b = sum
	}
}
