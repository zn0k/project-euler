package main

import (
	"fmt"
	"math/big"
)

func pow(a, b big.Int) big.Int {
	limit := b.Int64()
	mul := big.NewInt(a.Int64())
	for i := int64(1); i < limit; i++ {
		a = *a.Mul(&a, mul)
	}
	return a
}

func main() {
	tracker := make(map[string]bool)
	for a := int64(2); a <= 100; a++ {
		for b := int64(2); b <= 100; b++ {
			p := pow(*big.NewInt(a), *big.NewInt(b))
			tracker[p.String()] = true
		}
	}
	fmt.Printf("%d\n", len(tracker))
}
