package main

import (
	"fmt"
	"math/big"
)

func sharedDigits(n, d int64) (bool, int64, int64) {
	n1 := n / 10
	n2 := n % 10
	d1 := d / 10
	d2 := d % 10

	var n_keep, d_keep int64 = 0, 0
	var discard int64 = 0

	if n1 == d1 {
		n_keep, d_keep = n2, d2
		discard = n1
	} else if n1 == d2 {
		n_keep, d_keep = n2, d1
		discard = n1
	} else if n2 == d1 {
		n_keep, d_keep = n1, d2
		discard = n2
	} else if n2 == d2 {
		n_keep, d_keep = n1, d1
		discard = n2
	}
	if n_keep == 0 || d_keep == 0 || discard == 0 {
		return false, 0, 0
	}
	return true, n_keep, d_keep
}

func main() {
	var fractions []*big.Rat
	for n := int64(10); n < 100; n++ {
		for d := int64(10); d < 100; d++ {
			if n >= d {
				continue
			}
			ok, new_n, new_d := sharedDigits(n, d)
			if ok {
				a := big.NewRat(n, d)
				b := big.NewRat(new_n, new_d)
				if a.Cmp(b) == 0 {
					fractions = append(fractions, a)
				}
			}
		}
	}

	prod := big.NewRat(1, 1)
	for _, f := range fractions {
		prod = prod.Mul(prod, f)
	}
	fmt.Printf("%v\n", prod.Denom())
}
