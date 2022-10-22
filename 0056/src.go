package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func SumDigits(s string) int {
	sum := 0
	for _, r := range s {
		i, _ := strconv.ParseInt(string(r), 10, 32)
		sum += int(i)
	}
	return sum
}

func main() {
	max := 0
	for i := int64(1); i < 100; i++ {
		for j := int64(1); j < 100; j++ {
			a := big.NewInt(i)
			b := big.NewInt(j)
			a.Exp(a, b, nil)
			s := SumDigits(a.String())
			if s > max {
				max = s
			}
		}
	}
	fmt.Println(max)
}
