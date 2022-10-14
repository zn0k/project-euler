package main

import (
	"fmt"
	"math/big"
)

func fact(n big.Int) big.Int {
	i := big.NewInt(1)
	one := big.NewInt(1)
	for ; n.Cmp(one) == 1; n.Sub(&n, one) {
		i.Mul(i, &n)
	}
	return *i
}

func main() {
	f40 := fact(*big.NewInt(40))
	f20 := fact(*big.NewInt(20))
	f20.Mul(&f20, &f20)
	f40.Div(&f40, &f20)
	fmt.Printf("%s\n", f40.String())
}
