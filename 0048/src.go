package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	sum := big.NewInt(0)
	for i := int64(1); i <= 1000; i++ {
		exp := big.NewInt(i)
		var next big.Int
		next.Exp(exp, exp, nil)
		sum.Add(sum, &next)
	}

	cs := strings.Split(sum.Text(10), "")
	fmt.Println(strings.Join(cs[len(cs)-10:], ""))
}
