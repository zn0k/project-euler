package main

import (
	"fmt"
	"math/big"
)

func main() {
	total := 0
	for i := 0; i <= 100000; i++ {
		for p := 0; p <= 50; p++ {
			x := big.NewInt(int64(i))
			y := big.NewInt(int64(p))
			var n big.Int
			n.Exp(x, y, nil)
			if len(n.Text(10)) == p {
				total++
			}
		}
	}
	fmt.Println(total)
}
