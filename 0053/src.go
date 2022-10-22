package main

import (
	"fmt"
	"math/big"
)

type lookup map[int]*big.Int

func combinations(n, r int, l lookup) big.Int {
	var q, d big.Int
	q.Add(big.NewInt(0), l[n])
	d.Mul(l[r], l[n-r])
	q.Div(&q, &d)
	return q
}

func main() {
	factorials := lookup{0: big.NewInt(1), 1: big.NewInt(1)}

	for i := 2; i <= 100; i++ {
		n := big.NewInt(int64(i))
		n.Mul(n, factorials[i-1])
		factorials[i] = n
	}

	threshold := big.NewInt(1_000_000)
	count := 0
	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			c := combinations(n, r, factorials)
			if c.Cmp(threshold) > 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}
